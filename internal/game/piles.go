package game

// Tableau represents the seven columns of cards in the main playing area.
type Tableau struct {
	Columns [][]Card
}

// Stock represents the pile of cards that have not yet been dealt to the waste pile.
type Stock struct {
	Cards []Card
}

// Waste represents the pile of cards from the stock that can be played.
type Waste struct {
	Cards []Card
}

// Foundation represents the four piles where cards are moved to win the game.
type Foundation struct {
	Piles [][]Card
}
