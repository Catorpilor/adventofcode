package main

import "testing"

func TestHelper(t *testing.T) {
	st := []struct {
		name string
		str  string
		exp  int64
	}{
		{"testcase1", "2 * 3 + (4 * 5)", int64(26)},
		{"testcase2", "5 + (8 * 3 + 9 + 3 * 4 * 3)", int64(437)},
		{"testcase3", "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", int64(13632)},
		{"testcase4", "((5 * 6 + 7 + 4 + 3) * 9) * (4 + 8 * 9 + (5 + 2 + 2 + 4 * 7) + (9 * 7 + 4 * 5 + 3) * 2) + 6 * 6 * (7 * 2 + 5) + 7", int64(48485347)},
		{"testcaes5", "1 + (2 * 3) + (4 * (5 + 6))", int64(51)},
		{"testcase6", "(9 + (8 * 9 * 4 * 5 + 4) * 9 * 4 * 9 * 4) * ((6 * 9) + 3) * (6 * (3 * 5 + 6 + 9 + 2) * 6 + (4 + 5 + 3 + 2 + 2 * 4)) + 8 * 4 + 9", int64(522082381865)},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := helper(tt.str)
			if out != tt.exp {
				t.Fatalf("with input str:%s wanted %d but got %d", tt.str, tt.exp, out)
			}
		})
	}
}
