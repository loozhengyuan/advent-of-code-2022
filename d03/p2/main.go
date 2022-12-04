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

func getSharedItem(sacks ...string) rune {
	a := make(map[rune]int, 0)
	for _, sack := range sacks {
		t := make(map[rune]struct{}, len(sack))
		for _, item := range sack {
			// Only track item if not already seen in current sack
			if _, ok := t[item]; !ok {
				v, ok := a[item]
				if !ok {
					a[item] = 1
				} else {
					a[item] = v + 1
				}
			}
			t[item] = struct{}{}
		}
	}
	for k, v := range a {
		if v == len(sacks) {
			return k
		}
	}
	return 0
}

func processInput(r io.Reader) (int, error) {
	var l1, l2 string
	var p int
	s := bufio.NewScanner(r)
	for s.Scan() {
		if l1 == "" {
			l1 = s.Text()
			continue
		}
		if l2 == "" {
			l2 = s.Text()
			continue
		}
		p += getPriority(getSharedItem(l1, l2, s.Text()))
		l1 = ""
		l2 = ""
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
