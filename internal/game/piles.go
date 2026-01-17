package game

// Pile represents a stack of cards.
type Pile struct {
	Cards []*Card
}

// Push adds a card to the top of the pile.
func (p *Pile) Push(c *Card) {
	p.Cards = append(p.Cards, c)
}

// Pop removes and returns the top card from the pile. Returns nil if the pile is empty.
func (p *Pile) Pop() *Card {
	if len(p.Cards) == 0 {
		return nil
	}
	card := p.Cards[len(p.Cards)-1]
	p.Cards = p.Cards[:len(p.Cards)-1]
	return card
}

// Peek returns the top card of the pile without removing it. Returns nil if the pile is empty.
func (p *Pile) Peek() *Card {
	if len(p.Cards) == 0 {
		return nil
	}
	return p.Cards[len(p.Cards)-1]
}
