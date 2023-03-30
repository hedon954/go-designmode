package main

type Player struct {
	x      int
	y      int
	health int
}

func (p *Player) Move(direction string) {
	switch direction {
	case "up":
		p.y++
	case "down":
		p.y--
	case "left":
		p.x--
	case "right":
		p.x++
	}
}

func (p *Player) Attack(target *Player) {
	target.health -= 10
}
