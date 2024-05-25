package pet

import (
	"fmt"
	"time"

	"github.com/cassiusfive/gitpets/internal/gitstats"
)

type Pet struct {
	ownerGithub string
	dateCreated time.Time
	Name        string
	Level       int
	Xp          int
}

func Create(username, petname string) (Pet, error) {
	pet := Pet{
		ownerGithub: username,
		Name:        petname,
		dateCreated: time.Now(),
	}

	err := pet.SyncWithGit()
	if err != nil {
		return pet, err
	}

	return pet, nil
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
