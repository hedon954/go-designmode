package main

import (
	"fmt"
	"log"
)

// GameServerProxy is a proxy of GameServerImpl
type GameServerProxy struct {
	gameServer GameServer // service object
	username   string
	password   string
}

func (gsp *GameServerProxy) Connect() error {
	// check auth
	if gsp.username != "admin" || gsp.password != "password" {
		return fmt.Errorf("authentication failed")
	}
	if err := gsp.gameServer.Connect(); err != nil {
		return err
	}
	return nil
}

func (gsp *GameServerProxy) Disconnect() error {
	// clear cache
	gsp.username = ""
	gsp.password = ""
	if err := gsp.gameServer.Disconnect(); err != nil {
		return err
	}
	return nil
}

func (gsp *GameServerProxy) Send(message string) error {
	// log
	log.Print(gsp.username, " ", message)
	if err := gsp.gameServer.Send(message); err != nil {
		return err
	}
	return nil
}
