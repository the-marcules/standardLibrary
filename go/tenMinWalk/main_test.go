package main

import "testing"

var valid = [][]rune{
	[]rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'},
	[]rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e'},
	[]rune{'n', 's', 'e', 'w', 'w', 'w', 'e', 'e', 'n', 's'},
	[]rune{'n', 'e', 's', 'w', 's', 'w', 'n', 'e', 'n', 's'},
}

var invalid = [][]rune{
	[]rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e'},
	[]rune{'w'},
	[]rune{'n', 'n', 'n', 's', 'n', 's', 'n', 's', 'n', 's'},
	[]rune{'e', 'e', 'e', 'e', 'w', 'w', 's', 's', 's', 's'},
}

func TestIsValidWalk(t *testing.T) {
	for _, walk := range valid {
		t.Run("valid walks to be true", func(t *testing.T) {
			got := IsValidWalk(walk)

			if got != true {
				t.Error("Expected to be true but got false")
			}
		})
	}

	for _, walk := range invalid {
		t.Run("invalid walks to be false", func(t *testing.T) {
			got := IsValidWalk(walk)

			if got != false {
				t.Error("Expected to be false but got true")
			}
		})
	}
}
