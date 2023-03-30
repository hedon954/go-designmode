package main

type MoveCommand struct {
	player    *Player
	direction string
}

func (c *MoveCommand) Execute() {
	c.player.Move(c.direction)
}

type AttackCommand struct {
	palyer *Player
	target *Player
}

func (c *AttackCommand) Execute() {
	c.palyer.Attack(c.target)
}
