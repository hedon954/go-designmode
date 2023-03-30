package main

import (
	"fmt"
)

type GameServer struct {
	commands []Command
}

func (s *GameServer) AddCommands(commands ...Command) {
	s.commands = append(s.commands, commands...)
}

func (s *GameServer) ProcessCommands() {
	for _, c := range s.commands {
		c.Execute()
	}
}

func main() {
	server := GameServer{}

	player1 := &Player{}
	player2 := &Player{}

	moveCommand1 := &MoveCommand{player1, "right"}
	attackCommand1 := &AttackCommand{player1, player2}

	moveCommand2 := &MoveCommand{player2, "up"}
	attackCommand2 := &AttackCommand{player2, player1}

	server.AddCommands(moveCommand1, attackCommand1, moveCommand2, attackCommand2)

	server.ProcessCommands()

	fmt.Println(player1)
	fmt.Println(player2)
}
