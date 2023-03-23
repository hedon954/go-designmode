package main

import (
	"fmt"
)

func main() {
	player := &BasicPlayer{}

	// firePlayer decorates the basicPlayer
	firePlayer := &FirePlayer{player: player}
	firePlayer.ChooseClass("mage")
	skill := firePlayer.CastSkill()
	fmt.Println(skill)
}
