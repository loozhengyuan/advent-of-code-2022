package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getSectionRange(input string) (start, stop int, err error) {
	r := strings.Split(input, "-")
	if len(r) != 2 {
		return 0, 0, fmt.Errorf("failed to split range: %s", input)
	}
	startInt, err := strconv.Atoi(r[0])
	if err != nil {
		return 0, 0, fmt.Errorf("parse input: %#v: %w", r[0], err)
	}
	stopInt, err := strconv.Atoi(r[1])
	if err != nil {
		return 0, 0, fmt.Errorf("parse input: %#v: %w", r[1], err)
	}
	return startInt, stopInt, nil
}

func processInput(r io.Reader) (int, error) {
	var count int
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()

		// Derive start/stop ranges for elves
		pairs := strings.Split(line, ",")
		if len(pairs) != 2 {
			return 0, fmt.Errorf("failed to split line into 2 pairs: %s", line)
		}
		e1start, e1stop, err := getSectionRange(pairs[0])
		if err != nil {
			return 0, fmt.Errorf("get elf 1 range: %s", pairs[0])
		}
		e2start, e2stop, err := getSectionRange(pairs[1])
		if err != nil {
			return 0, fmt.Errorf("get elf 2 range: %s", pairs[1])
		}

		// Determine if either elves overlap with one another
		if e1start >= e2start && e1stop <= e2stop {
			count++
			continue
		}
		if e2start >= e1start && e2stop <= e1stop {
			count++
			continue
		}
	}
	if err := s.Err(); err != nil {
		return 0, fmt.Errorf("scan input: %w", err)
	}
	return count, nil
}

func run() error {
	f, err := os.Open("d04/actual.txt")
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
