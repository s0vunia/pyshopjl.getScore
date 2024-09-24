package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const TIMESTAMPS_COUNT = 50000

const PROBABILITY_SCORE_CHANGED = 0.0001

const PROBABILITY_HOME_SCORE = 0.55

const OFFSET_MAX_STEP = 3

type Score struct {
	Home int
	Away int
}

type ScoreStamp struct {
	Offset int
	Score  Score
}

func main() {
	var stamps = fillScores()

	for _, stamp := range *stamps {
		fmt.Printf("%v: %v -- %v\n", stamp.Offset, stamp.Score.Home, stamp.Score.Away)
	}

	testOffset := 100000
	result := getScore(*stamps, testOffset)
	fmt.Printf("Score at offset %d: %v -- %v\n", testOffset, result.Home, result.Away)
}

func generateStamp(previousValue ScoreStamp) ScoreStamp {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	scoreChanged := random.Float32() > 1-PROBABILITY_SCORE_CHANGED
	homeScoreChange := 0
	if scoreChanged && random.Float32() > 1-PROBABILITY_HOME_SCORE {
		homeScoreChange = 1
	}

	awayScoreChange := 0
	if scoreChanged && homeScoreChange == 0 {
		awayScoreChange = 1
	}

	offsetChange := int(math.Floor(random.Float64()*OFFSET_MAX_STEP)) + 1

	return ScoreStamp{
		Offset: previousValue.Offset + offsetChange,
		Score: Score{
			Home: previousValue.Score.Home + homeScoreChange,
			Away: previousValue.Score.Away + awayScoreChange,
		},
	}
}

func fillScores() *[]ScoreStamp {

	scores := make([]ScoreStamp, TIMESTAMPS_COUNT)
	prevScore := ScoreStamp{
		Offset: 0,
		Score:  Score{Home: 0, Away: 0},
	}
	scores[0] = prevScore

	for i := 1; i < TIMESTAMPS_COUNT; i++ {
		scores[i] = generateStamp(prevScore)
		prevScore = scores[i]
	}

	return &scores
}

// getScore returns the score at a given offset in the game.
// It handles cases where the exact offset may not exist in the game stamps.
//
// Parameters:
//   - gameStamps: A slice of ScoreStamp representing the game's score history.
//   - offset: An integer representing the time offset to query.
//
// Returns:
//
//	A Score struct containing the home and away scores at the given offset.
//
// Time complexity: O(log n), Space complexity: O(1), where n is the number of game stamps.
func getScore(gameStamps []ScoreStamp, offset int) Score {
	if len(gameStamps) == 0 {
		return Score{Home: 0, Away: 0}
	}

	if offset < gameStamps[0].Offset {
		return Score{Home: 0, Away: 0}
	}

	if offset >= gameStamps[len(gameStamps)-1].Offset {
		return gameStamps[len(gameStamps)-1].Score
	}

	left, right := 0, len(gameStamps)-1
	for left < right {
		mid := left + (right-left+1)/2
		if gameStamps[mid].Offset <= offset {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return gameStamps[left].Score
}
