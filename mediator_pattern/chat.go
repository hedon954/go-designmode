package main

import (
	"fmt"
)

// Mediator is the chat room mediator
type Mediator interface {
	sendMessage(user User, message string)
}

// User is a user of the chat room,
// it holds the mediator and communicates with it
type User struct {
	name     string
	mediator Mediator
}

// NewUser creates a new user
func NewUser(name string, mediator Mediator) *User {
	return &User{
		name:     name,
		mediator: mediator,
	}
}

// SendMessage sends a message
// it just needs to send msg to the mediator
// needs not to cara about how many users there are
func (u *User) SendMessage(msg string) {
	u.mediator.sendMessage(*u, msg)
}

// ChatRoom is the chat room, it implements the Mediator
type ChatRoom struct {
	users []*User
}

func (c *ChatRoom) sendMessage(user User, message string) {
	for _, u := range c.users {
		if u.name != user.name {
			fmt.Printf("%s recevied the message [%s] from %s\n", u.name, message, user.name)
		}
	}
}

func (c *ChatRoom) AddUsers(users ...*User) {
	c.users = append(c.users, users...)
}

func (c *ChatRoom) RemoveUsers(names ...string) {
	for _, name := range names {
		for i, u := range c.users {
			if u.name == name {
				c.users = append(c.users[:i], c.users[i+1:]...)
				break
			}
		}
	}
}

func main() {
	room := &ChatRoom{}
	u1 := &User{"Alice", room}
	u2 := &User{"Bob", room}
	u3 := &User{"Charlie", room}
	room.AddUsers(u1, u2, u3)

	u1.SendMessage("hello")
	// Bob recevied the message [hello] from Alice
	// Charlie recevied the message [hello] from Alice

	u2.SendMessage("hi")
	// Alice recevied the message [hi] from Bob
	// Charlie recevied the message [hi] from Bob

	u3.SendMessage("here")
	// Alice recevied the message [here] from Charlie
	// Bob recevied the message [here] from Charlie
}
