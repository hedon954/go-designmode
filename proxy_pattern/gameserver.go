package main

// GameServer defines the interface
type GameServer interface {
	Connect() error        // connect to the game server
	Disconnect() error     // disconnect from the game server
	Send(msg string) error // send message to the game server
}

// GameServerImpl is a implements of the GameServer
type GameServerImpl struct {
	address string
}

func (gs *GameServerImpl) Connect() error {
	// build the connection with game server
	return nil
}

func (gs *GameServerImpl) Disconnect() error {
	// cut the connection with game server
	return nil
}

func (gs *GameServerImpl) Send(message string) error {
	// send message to the game server
	return nil
}
