package main

import (
	"os"
	"testing"
)

func TestProcessInput(t *testing.T) {
	cases := map[string]struct {
		filename string
		want     int
	}{
		"example": {
			filename: "../example.txt",
			want:     4,
		},
		"actual": {
			filename: "../actual.txt",
			want:     903,
		},
	}

	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			f, err := os.Open(tc.filename)
			if err != nil {
				t.Fatalf("failed to open file: %v", err)
			}
			defer f.Close()

			got, err := processInput(f)
			if err != nil {
				t.Fatalf("failed to invoke function: %v", err)
			}
			if got != tc.want {
				t.Errorf("output mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.want)
			}
		})
	}
}
