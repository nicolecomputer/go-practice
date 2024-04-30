package day02

type CubeConfiguration struct {
	Red   int
	Blue  int
	Green int
}

func ConfigurationIsPossible(configuration CubeConfiguration, game *Game) bool {
	for _, set := range game.RevealedSet {
		for _, cube := range set.Seen {
			if cube.Color == "red" && *cube.Count > configuration.Red {
				return false
			} else if cube.Color == "blue" && *cube.Count > configuration.Blue {
				return false
			} else if cube.Color == "green" && *cube.Count > configuration.Green {
				return false
			}
		}
	}
	return true
}

func TotalValueForPossibleGames(configuration CubeConfiguration, gameRecord *GameRecord) int {
	total := 0

	for _, game := range gameRecord.Games {
		if ConfigurationIsPossible(configuration, game) {
			total += *game.Id
		}
	}

	return total
}
