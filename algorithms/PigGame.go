package algorithms

import "math/rand"

const (
	win            = 100
	gamesPerSeries = 10
)

// Implements the Pig dice game - https://en.wikipedia.org/wiki/Pig_(dice_game)

// The following structure holds the total scores and the running score
type scores struct {
	player   int // This player's score
	opponent int // Opponent's score
	thisTurn int // Score for this turn
}

// action is a function type. The function must take a scores parameters and returns
// multiple values: the resulting scores and whether the current turn is over
type action func(current scores) (result scores, turnsIsOver bool)

// The following two functions can be assigned to the 'action' function type
// This function rolls the dice and either increments the player's turn score
// or resets it to zero and swaps scores
func roll(current scores) (result scores, turnIsOver bool) {
	// Roll a 6-faced dice to get a number from 1 to 6 inclusive
	var diceRoll = rand.Intn(6) + 1

	// If dice roll is not 1, add the dice value to the current score
	if diceRoll != 1 {
		return scores{
			player:   current.player,
			opponent: current.opponent,
			thisTurn: current.thisTurn + diceRoll,
		}, false
	}

	// Otherwise, scores for this turn are set to zero and roles swapped
	return scores{
		player:   current.opponent,
		opponent: current.player,
		thisTurn: 0,
	}, true
}

// This function adds this turn's score to the player's total score and swaps the scores
func hold(current scores) (scores, bool) {
	return scores{
		player:   current.opponent,
		opponent: current.player + current.thisTurn,
		thisTurn: 0,
	}, true
}

// strategy is a function type: it takes a scores structure and returns a function
// The function that is returned is of type action
type strategy func(scores) action

// Simulate a game of Pig by calling an action to update the score until one player reaches
//100 points. Each action is selected by calling the strategy function associated with the
//current player
func rollButStayAtK(k int) strategy {
	return func(s scores) action {
		if s.thisTurn >= k {
			return hold
		}
		return roll
	}
}

func Play() (winnerNumber int) {
	// Strategies for both players: Player 1 keeps on rolling, but holds on 50,
	// while player 2 keeps on rolling but holds on 20
	var p1Strategy = rollButStayAtK(50)
	var p2Strategy = rollButStayAtK(20)
	var strategies = []strategy{p1Strategy, p2Strategy}

	var s scores
	var turnIsOver bool

	// Let player 1 start
	var currentPlayer = 0

	// Keep on playing while scores for both players are under the required win value
	for s.player+s.thisTurn < win {

		// Get strategy and actin for this player
		var currentPlayerStrategy strategy = strategies[currentPlayer]
		var currentPlayerAction action = currentPlayerStrategy(s)

		// Roll the dice
		s, turnIsOver = currentPlayerAction(s)

		// Switch to the other player if turn is over, otherwise roll dice again
		if turnIsOver {
			currentPlayer = (currentPlayer + 1) % 2 // switch between 0 and 1
		}
	}

	return currentPlayer // return which player won
}
