package day02

import (
	"testing"
)

func TestConfigurationIsPossible(t *testing.T) {
	samples := []struct {
		game     string
		possible bool
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true},
		{"Game 6: 0 red, 14 blue, 0 green; 2 blue, 1 red, 2 green", false},
		{"Game 6: 13 red, 0 blue, 0 green; 2 blue, 1 red, 2 green", false},
		{"Game 6: 0 red, 0 blue, 15 green; 2 blue, 1 red, 2 green", false},
		{"Game 6: 12 red, 13 blue, 14 green; 2 blue, 1 red, 2 green", true},
	}

	for _, test := range samples {
		t.Run(test.game, func(t *testing.T) {
			game := ParseGame(test.game)

			possible := ConfigurationIsPossible(CubeConfiguration{12, 13, 14}, game)

			if possible != test.possible {
				t.Errorf("Expected score %t receieved %t", test.possible, possible)
			}
		})
	}
}

func TestTotalValueForPossibleGames(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	games := ParseGameRecord(input)

	want := 8
	got := TotalValueForPossibleGames(CubeConfiguration{12, 13, 14}, games)

	if want != got {
		t.Errorf("Expected %d got %d", want, got)
	}
}
