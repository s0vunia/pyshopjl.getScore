# Game Score Tracker
A Go package for tracking and retrieving game scores at specific time offsets.

## Features
- Efficient score retrieval using binary search
- Handling of various game scenarios (empty game, single score, multiple scores)
- Comprehensive test suite with high code coverage

## Requirements
- Go 1.22.6 or later

## Installation
1. Clone the repository:
   ```
   git clone https://github.com/s0vunia/pyshopjl.getScore.git
   cd pyshopjl.getScore
   ```

2. Import the package in your Go project:
   ```go
   import "github.com/s0vunia/pyshopjl.getScore"
   ```

## Usage
Here's a basic example of how to use the `getScore` function:

```go
gameStamps := []ScoreStamp{
    {Offset: 0, Score: Score{Home: 0, Away: 0}},
    {Offset: 10, Score: Score{Home: 1, Away: 0}},
    {Offset: 20, Score: Score{Home: 1, Away: 1}},
}

score := getScore(gameStamps, 15)
fmt.Printf("Score at offset 15: Home %d - Away %d\n", score.Home, score.Away)
```

## How It Works
1. The `getScore` function takes a slice of `ScoreStamp` and an offset.
2. It uses binary search to efficiently find the appropriate score for the given offset.
3. Edge cases like empty input, out-of-range offsets, and exact matches are handled.

## Testing
Run the tests with the following command:
```
go test -v
```

To check test coverage:
```
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.