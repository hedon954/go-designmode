package main

import (
	"fmt"
)

func main() {
	gameServer := &GameServerImpl{address: "localhost:12345"}
	gameServerProxy := &GameServerProxy{
		gameServer: gameServer,
		username:   "admin",
		password:   "password",
	}
	if err := gameServerProxy.Connect(); err != nil {
		fmt.Printf("error connecting to game server: %v", err)
		return
	}
	defer gameServerProxy.Disconnect()

	if err := gameServerProxy.Send("hello"); err != nil {
		fmt.Printf("error sending message to game server: %v", err)
		return
	}
}
