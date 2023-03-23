package main

// Player is the Component
type Player interface {
	ChooseClass(class string)
	CastSkill() string
}

// BasicPlayer is the concrete who implements the Player interface
type BasicPlayer struct {
	Class string
}

func (c *BasicPlayer) ChooseClass(class string) {
	c.Class = class
}

func (c *BasicPlayer) CastSkill() string {
	return "basic class"
}

// FirePlayer is the decorator of Player
type FirePlayer struct {
	player Player
}

func (f *FirePlayer) ChooseClass(class string) {
	f.player.ChooseClass(class)
}

func (f *FirePlayer) CastSkill() string {
	return f.player.CastSkill() + ", fireball"
}
