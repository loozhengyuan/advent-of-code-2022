package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func processInput(r io.Reader) (int, error) {
	var max, cur int
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()

		// Check max count and reset if empty line
		if line == "" {
			if cur > max {
				max = cur
			}
			cur = 0
			continue
		}

		// Convert to integer to add value to current counter
		val, err := strconv.Atoi(s.Text())
		if err != nil {
			return 0, fmt.Errorf("parse input: %#v: %w", val, err)
		}
		cur += val
	}
	if err := s.Err(); err != nil {
		return 0, fmt.Errorf("scan input: %w", err)
	}

	// Re-check max count for the last chunk
	if cur > max {
		max = cur
	}
	return max, nil
}

func run() error {
	f, err := os.Open("d01/actual.txt")
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
