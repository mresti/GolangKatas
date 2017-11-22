package tennis

import (
	"fmt"
	"math"
	"testing"
)

func TestTennisGame(t *testing.T) {
	var tests = []struct {
		player1Score  int
		player2Score  int
		expectedScore string
	}{
		{player1Score: 0, player2Score: 0, expectedScore: "Love-All"},
		{player1Score: 1, player2Score: 1, expectedScore: "Fifteen-All"},
		{player1Score: 2, player2Score: 2, expectedScore: "Thirty-All"},
		{player1Score: 3, player2Score: 3, expectedScore: "Deuce"},
		{player1Score: 4, player2Score: 4, expectedScore: "Deuce"},

		{player1Score: 1, player2Score: 0, expectedScore: "Fifteen-Love"},
		{player1Score: 0, player2Score: 1, expectedScore: "Love-Fifteen"},
		{player1Score: 2, player2Score: 0, expectedScore: "Thirty-Love"},
		{player1Score: 0, player2Score: 2, expectedScore: "Love-Thirty"},
		{player1Score: 3, player2Score: 0, expectedScore: "Forty-Love"},
		{player1Score: 0, player2Score: 3, expectedScore: "Love-Forty"},
		{player1Score: 4, player2Score: 0, expectedScore: "Win for player1"},
		{player1Score: 0, player2Score: 4, expectedScore: "Win for player2"},

		{player1Score: 2, player2Score: 1, expectedScore: "Thirty-Fifteen"},
		{player1Score: 1, player2Score: 2, expectedScore: "Fifteen-Thirty"},
		{player1Score: 3, player2Score: 1, expectedScore: "Forty-Fifteen"},
		{player1Score: 1, player2Score: 3, expectedScore: "Fifteen-Forty"},
		{player1Score: 4, player2Score: 1, expectedScore: "Win for player1"},
		{player1Score: 1, player2Score: 4, expectedScore: "Win for player2"},

		{player1Score: 3, player2Score: 2, expectedScore: "Forty-Thirty"},
		{player1Score: 2, player2Score: 3, expectedScore: "Thirty-Forty"},
		{player1Score: 4, player2Score: 2, expectedScore: "Win for player1"},
		{player1Score: 2, player2Score: 4, expectedScore: "Win for player2"},

		{player1Score: 4, player2Score: 3, expectedScore: "Advantage player1"},
		{player1Score: 3, player2Score: 4, expectedScore: "Advantage player2"},
		{player1Score: 5, player2Score: 4, expectedScore: "Advantage player1"},
		{player1Score: 4, player2Score: 5, expectedScore: "Advantage player2"},
		{player1Score: 15, player2Score: 14, expectedScore: "Advantage player1"},
		{player1Score: 14, player2Score: 15, expectedScore: "Advantage player2"},

		{player1Score: 6, player2Score: 4, expectedScore: "Win for player1"},
		{player1Score: 4, player2Score: 6, expectedScore: "Win for player2"},
		{player1Score: 16, player2Score: 14, expectedScore: "Win for player1"},
		{player1Score: 14, player2Score: 16, expectedScore: "Win for player2"},
	}
	for _, test := range tests {
		inputMessage := fmt.Sprintf("P1=%v,P2=%v", test.player1Score, test.player2Score)
		t.Run(inputMessage, func(t *testing.T) {
			tennis := MatchTennis("player1", "player2")
			highScore := int(math.Max(float64(test.player1Score), float64(test.player2Score)))
			for i := 0; i < highScore; i++ {
				if i < test.player1Score {
					tennis.WonPoint("player1")
				}
				if i < test.player2Score {
					tennis.WonPoint("player2")
				}
			}

			result := tennis.GetScore()
			if result != test.expectedScore {
				t.Fatalf("P1:%v, P2:%v returned %v, want %v", test.player1Score, test.player2Score, result, test.expectedScore)
			}
		})
	}
}
