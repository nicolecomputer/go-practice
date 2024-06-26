package day04

import "testing"

func TestScore(t *testing.T) {
	scoreTests := []struct {
		game  string
		score int
	}{
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 8},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 2},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 1},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 0},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 0},
	}

	for _, test := range scoreTests {
		t.Run(test.game, func(t *testing.T) {
			game := ParseGame(test.game)

			score := Score(game)

			if score != test.score {
				t.Errorf("Expected score %d receieved %d", test.score, score)
			}
		})
	}
}

func TestTotalScore(t *testing.T) {
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
	Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
	Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
	Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
	Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
	games := ParseGames(input)

	score := TotalScore(games)
	want := 13

	if score != want {
		t.Errorf("Expected score %d receieved %d", want, score)
	}
}
