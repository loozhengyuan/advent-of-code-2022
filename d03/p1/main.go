package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func getPriority(r rune) int {
	if unicode.IsLower(r) {
		return int(r) - 96
	}
	if unicode.IsUpper(r) {
		return int(r) - 38
	}
	return 0
}

func getSharedItem(contents string) rune {
	m := make(map[rune]struct{}, 0)
	max := len(contents)
	for i, c := range contents {
		if i < max/2 {
			if _, ok := m[c]; !ok {
				m[c] = struct{}{}
			}
		} else {
			if _, ok := m[c]; ok {
				return c
			}
		}
	}
	return 0
}

func processInput(r io.Reader) (int, error) {
	var p int
	s := bufio.NewScanner(r)
	for s.Scan() {
		p += getPriority(getSharedItem(s.Text()))
	}
	if err := s.Err(); err != nil {
		return 0, fmt.Errorf("scan input: %w", err)
	}
	return p, nil
}

func run() error {
	f, err := os.Open("d03/actual.txt")
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
