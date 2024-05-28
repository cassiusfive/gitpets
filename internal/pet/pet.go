package pet

import (
	"errors"
	"fmt"
	"hash/fnv"
	"math/rand"
	"strconv"
	"time"

	"github.com/cassiusfive/gitpets/internal/gitstats"
)

var validSpecies = [...]string{"fox", "wolf"}
var moods = [...]string{"happy", "sad", "locked in", "relaxed", "bored", "sleepy"}

type Pet struct {
	ownerGithub string
	dateCreated time.Time
	Name        string
	Species     string
	Level       int
	Xp          int
	Mood        string
}

func Create(username, petname, species string) (Pet, error) {
	pet := Pet{
		ownerGithub: username,
		Name:        petname,
		Species:     species,
		Mood:        randomMood(username, petname),
		dateCreated: time.Now(),
	}

	if !isValidSpecies(species) {
		return pet, errors.New("Invalid species")
	}

	err := pet.SyncWithGit()
	if err != nil {
		return pet, err
	}

	return pet, nil
}

func isValidSpecies(species string) bool {
	for _, s := range validSpecies {
		if species == s {
			return true
		}
	}
	return false
}

func randomMood(username, petname string) string {
	hash := fnv.New64()
	hash.Write([]byte(username))
	hash.Write([]byte(petname))
	hash.Write([]byte(strconv.Itoa(int(time.Now().Unix() / 4000))))
	s := rand.NewSource(int64(hash.Sum64()))
	r := rand.New(s)
	return moods[r.Intn(len(moods))]
}

func ExperienceToLevel(level int) int {
	return level*level + 4
}

func calculateExperience(stats gitstats.GitStats) (level, xpRemainder int) {
	totalXp := stats.TotalCommits*2 + stats.MergedPRs*10 + stats.ContributedTo*20
	level = 1
	for totalXp >= ExperienceToLevel(level) {
		totalXp -= ExperienceToLevel(level)
		level++
	}
	return level, totalXp
}

func (pet *Pet) Age() string {
	age := time.Now().Sub(pet.dateCreated)
	years := int(age.Hours()) / (24 * 365)
	months := int(age.Hours()) / (24 * 30)
	days := int(age.Hours()) / (24)
	hours := int(age.Hours())
	minutes := int(age.Minutes())
	if years >= 1 {
		if years > 1 {
			return fmt.Sprintf("%d years", years)
		}
		return fmt.Sprintf("%d year", years)
	} else if months >= 1 {
		if years > 1 {
			return fmt.Sprintf("%d months", months)
		}
		return fmt.Sprintf("%d month", months)
	} else if days >= 1 {
		if years > 1 {
			return fmt.Sprintf("%d years", years)
		}
		return fmt.Sprintf("%d year", years)
	} else if hours >= 1 {
		if hours > 1 {
			return fmt.Sprintf("%d hours", hours)
		}
		return fmt.Sprintf("%d hour", hours)
	}
	return fmt.Sprintf("%d minutes", minutes)
}

func (pet *Pet) SyncWithGit() error {
	stats, err := gitstats.GetStats(pet.ownerGithub)
	if err != nil {
		return err
	}

	pet.Level, pet.Xp = calculateExperience(stats)
	return nil
}
