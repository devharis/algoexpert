package easy

//
// Link: https://www.algoexpert.io/questions/Tournament%20Winner
//
//
// Time complexity O(n)
// Space complexity O(k)

// input: competitions= [
//   ["HTML", "C#"],
//   ["C#", "Python"],
//   ["Python", "HTML"],
// ]
// output: "Python"

const HOME_TEAM_WON = 1
const HOME_TEAM = 0
const AWAY_TEAM = 1

func TournamentWinner(competitions [][]string, results []int) string {

	scores := make(map[string]int)
	leadingTeam := ""
	leadingScore := 0

	for i, competition := range competitions {
		home := competition[HOME_TEAM]
		away := competition[AWAY_TEAM]

		result := results[i]

		if result == HOME_TEAM_WON {
			// HOME TEAM WON
			scores = calculateScore(home, scores)
			if scores[home] > leadingScore {
				leadingTeam = home
				leadingScore = scores[home]
			}
		} else {
			// AWAY TEAM WON
			scores = calculateScore(away, scores)
			if scores[away] > leadingScore {
				leadingTeam = away
				leadingScore = scores[away]
			}
		}
	}

	return leadingTeam
}

func calculateScore(team string, scores map[string]int) map[string]int {
	score := 0
	if value, exists := scores[team]; exists {
		score = value + 3
	} else {
		score = 3
	}
	scores[team] = score

	return scores
}
