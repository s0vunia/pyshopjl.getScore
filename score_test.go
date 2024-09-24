package main

import (
	"testing"
)

type TestCase struct {
	name       string
	gameStamps []ScoreStamp
	offset     int
	expected   Score
}

func assertScore(t *testing.T, got, want Score) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func runTestCases(t *testing.T, testCases []TestCase) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := getScore(tc.gameStamps, tc.offset)
			assertScore(t, got, tc.expected)
		})
	}
}

func TestGetScore(t *testing.T) {
	gameStamps := []ScoreStamp{
		{Offset: 0, Score: Score{Home: 0, Away: 0}},
		{Offset: 10, Score: Score{Home: 1, Away: 0}},
		{Offset: 20, Score: Score{Home: 1, Away: 1}},
		{Offset: 30, Score: Score{Home: 2, Away: 1}},
		{Offset: 40, Score: Score{Home: 2, Away: 2}},
	}

	testCases := []TestCase{
		{"ExactMatchStart", gameStamps, 0, Score{Home: 0, Away: 0}},
		{"ExactMatchMiddle", gameStamps, 20, Score{Home: 1, Away: 1}},
		{"ExactMatchEnd", gameStamps, 40, Score{Home: 2, Away: 2}},
		{"OffsetBeforeFirstStamp", gameStamps, -5, Score{Home: 0, Away: 0}},
		{"OffsetAfterLastStamp", gameStamps, 50, Score{Home: 2, Away: 2}},
		{"OffsetBetweenStamps", gameStamps, 15, Score{Home: 1, Away: 0}},
		{"OffsetAtUpperBoundary", gameStamps, 39, Score{Home: 2, Away: 1}},
	}

	runTestCases(t, testCases)
}

func TestGetScoreEdgeCases(t *testing.T) {
	testCases := []TestCase{
		{"EmptyGameStamps", []ScoreStamp{}, 10, Score{Home: 0, Away: 0}},
		{"SingleStampBefore", []ScoreStamp{{Offset: 5, Score: Score{Home: 1, Away: 0}}}, 0, Score{Home: 0, Away: 0}},
		{"SingleStampAt", []ScoreStamp{{Offset: 5, Score: Score{Home: 1, Away: 0}}}, 5, Score{Home: 1, Away: 0}},
		{"SingleStampAfter", []ScoreStamp{{Offset: 5, Score: Score{Home: 1, Away: 0}}}, 10, Score{Home: 1, Away: 0}},
	}

	runTestCases(t, testCases)
}

func TestGetScoreLargeDataset(t *testing.T) {
	largeGameStamps := make([]ScoreStamp, 10000)
	for i := 0; i < 10000; i++ {
		largeGameStamps[i] = ScoreStamp{Offset: i * 10, Score: Score{Home: i, Away: i}}
	}

	testCases := []TestCase{
		{"LargeDatasetStart", largeGameStamps, 0, Score{Home: 0, Away: 0}},
		{"LargeDatasetMiddle", largeGameStamps, 50000, Score{Home: 5000, Away: 5000}},
		{"LargeDatasetEnd", largeGameStamps, 99990, Score{Home: 9999, Away: 9999}},
	}

	runTestCases(t, testCases)
}
