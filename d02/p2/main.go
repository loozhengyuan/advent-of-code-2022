package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type OpponentHand string

const (
	OpponentHandUnspecified = ""
	OpponentHandRock        = "A"
	OpponentHandPaper       = "B"
	OpponentHandScissors    = "C"
)

type PlayerHand string

const (
	PlayerHandUnspecified = ""
	PlayerHandRock        = "X"
	PlayerHandPaper       = "Y"
	PlayerHandScissors    = "Z"
)

type GameResult string

const (
	GameResultUnspecified = ""
	GameResultLose        = "X"
	GameResultDraw        = "Y"
	GameResultWin         = "Z"
)

func getGameScore(playerHand, opponentHand string) int {
	if (playerHand == PlayerHandRock && opponentHand == OpponentHandRock) ||
		(playerHand == PlayerHandPaper && opponentHand == OpponentHandPaper) ||
		(playerHand == PlayerHandScissors && opponentHand == OpponentHandScissors) {
		return 3
	}
	if (playerHand == PlayerHandRock && opponentHand == OpponentHandScissors) ||
		(playerHand == PlayerHandPaper && opponentHand == OpponentHandRock) ||
		(playerHand == PlayerHandScissors && opponentHand == OpponentHandPaper) {
		return 6
	}
	return 0
}

func getHandScore(playerHand string) int {
	if playerHand == PlayerHandRock {
		return 1
	}
	if playerHand == PlayerHandPaper {
		return 2
	}
	if playerHand == PlayerHandScissors {
		return 3
	}
	return 0
}

func getPlayerHandFromResult(opponentHand, result string) string {
	if (opponentHand == OpponentHandRock && result == GameResultLose) ||
		(opponentHand == OpponentHandScissors && result == GameResultDraw) ||
		(opponentHand == OpponentHandPaper && result == GameResultWin) {
		return PlayerHandScissors
	}
	if (opponentHand == OpponentHandPaper && result == GameResultLose) ||
		(opponentHand == OpponentHandRock && result == GameResultDraw) ||
		(opponentHand == OpponentHandScissors && result == GameResultWin) {
		return PlayerHandRock
	}
	if (opponentHand == OpponentHandScissors && result == GameResultLose) ||
		(opponentHand == OpponentHandPaper && result == GameResultDraw) ||
		(opponentHand == OpponentHandRock && result == GameResultWin) {
		return PlayerHandPaper
	}
	return PlayerHandUnspecified
}

func processInput(r io.Reader) (int, error) {
	var score int
	s := bufio.NewScanner(r)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		oh, res := fields[0], fields[1]
		ph := getPlayerHandFromResult(oh, res)
		score += getGameScore(ph, oh)
		score += getHandScore(ph)
	}
	if err := s.Err(); err != nil {
		return 0, fmt.Errorf("scan input: %w", err)
	}
	return score, nil
}

func run() error {
	f, err := os.Open("d02/actual.txt")
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	r, err := processInput(f)
	if err != nil {
		return fmt.Errorf("process input: %w", err)
	}

	fmt.Fprintf(os.Stdout, "%d\n", r)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
