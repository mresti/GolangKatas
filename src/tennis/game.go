package tennis

// Game define methods by a tennis game
type Game interface {
	WonPoint(playerName string)
	GetScore() string
}

// tennisGame represent a Tennis Game
type tennisGame struct {
	score1      int
	score2      int
	player1Name string
	player2Name string
}

// MatchTennis return a New tennisGame
func MatchTennis(p1Name string, p2Name string) Game {
	return &tennisGame{player1Name: p1Name, player2Name: p2Name}
}

func (game *tennisGame) WonPoint(playerName string) {
	if playerName == "player1" {
		game.score1++
	} else {
		game.score2++
	}
}

func (game *tennisGame) GetScore() string {
	score := ""
	if game.score1 < 4 && game.score2 < 4 && !(game.score1+game.score2 == 6) {
		puntuation := []string{"Love", "Fifteen", "Thirty", "Forty"}
		score = puntuation[game.score1] + "-"

		if game.score1 == game.score2 {
			return score + "All"
		}
		return score + puntuation[game.score2]
	}

	if game.score1 == game.score2 {
		return "Deuce"
	}
	if game.score1 > game.score2 {
		score = game.player1Name
	} else {
		score = game.player2Name
	}
	advantage := game.score1 - game.score2
	if (advantage * advantage) == 1 {
		return "Advantage " + score
	}

	return "Win for " + score
}
