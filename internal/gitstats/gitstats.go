package gitstats

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/machinebox/graphql"
)

type GitStats struct {
	name          string
	totalCommits  int
	contributedTo int
	mergedPRs     int
}

type QueryStatsResponse struct {
	User struct {
		Name                    string
		ContributionsCollection struct {
			TotalCommitContributions int
			ReposContributedTo       int
			PullRequestContributions struct {
				TotalCount int
			}
		}
		MergedPRs struct {
			totalCount int
		}
	}
}

func GetStats(username string) (GitStats, error) {
	client := graphql.NewClient("https://api.github.com/graphql")

	query := fmt.Sprintf(`
		query {
		  user(login: "%s") {
		    name
		    contributionsCollection {
		      totalCommitContributions
					reposContributedTo: totalRepositoriesWithContributedCommits
		      pullRequestContributions {
		       totalCount
		      }
		    }
		    mergedPRs: pullRequests(states: MERGED) {
		      totalCount
		    }
		  }
		}
	`, username)

	request := graphql.NewRequest(query)
	request.Header.Add("Authorization", "Bearer "+os.Getenv("GITHUB_TOKEN"))
	var response QueryStatsResponse
	err := client.Run(context.Background(), request, &response)
	if err != nil {
		if strings.HasPrefix(err.Error(), "graphql: Could not resolve to a User") {
			return GitStats{name: "", totalCommits: 0, contributedTo: 0, mergedPRs: 0}, errors.New("user not found")
		}
	}

	return GitStats{
		name:          response.User.Name,
		totalCommits:  response.User.ContributionsCollection.TotalCommitContributions,
		contributedTo: response.User.ContributionsCollection.ReposContributedTo,
		mergedPRs:     response.User.MergedPRs.totalCount,
	}, nil
}
