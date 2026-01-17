package game

import "strconv"

// Suit represents the suit of a playing card.
type Suit int

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

// String returns the string representation of a suit.
func (s Suit) String() string {
	switch s {
	case Spades:
		return "♠"
	case Hearts:
		return "♥"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣"
	default:
		return ""
	}
}

// Color returns "Red" or "Black".
func (s Suit) Color() string {
	if s == Hearts || s == Diamonds {
		return "Red"
	}
	return "Black"
}

// Rank represents the rank of a playing card.
type Rank int

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// String returns the string representation of a rank.
func (r Rank) String() string {
	switch r {
	case Ace:
		return "A"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case 10:
		return "10"
	default:
		return strconv.Itoa(int(r))
	}
}

// Card represents a single playing card.
type Card struct {
	Suit   Suit
	Rank   Rank
	FaceUp bool
}
