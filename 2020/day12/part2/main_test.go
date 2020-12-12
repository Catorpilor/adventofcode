package main

import "testing"

// func initWithTesting(t *testing.T) map[string]int {
// 	wp := map[string]int{
// 		"N": 1,
// 		"E": 10,
// 		"S": 0,
// 		"W": 0,
// 	}
// 	t.Cleanup(func() {
// 		for k := range wp {
// 			delete(wp, k)
// 		}
// 	})
// 	return wp
// }

func TestDir(t *testing.T) {
	// wp := initWithTesting(t)
	st := []struct {
		name   string
		action byte
		unit   int
		exp    string
	}{
		{"L90", 'L', 90, "NW"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			calc(tt.action, tt.unit)
			out := dir(waypoint)
			if out != tt.exp {
				t.Fatalf("wanted %s but got %s", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
