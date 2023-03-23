package main

import (
	"fmt"
	"time"
)

// Action is the composite component
type Action interface {

	// Execute does the action
	Execute()
}

// PlayerAction is the composite leaf, it implements the Action interface
type PlayerAction struct {
	playerId   int
	actionTime time.Time
}

func (p *PlayerAction) Execute() {
	fmt.Printf("Player %d performed actions at %v\n", p.playerId, p.actionTime)
}

// CompositeAction is the composite, it is the collection of Action
type CompositeAction struct {
	actions []Action
}

func (c *CompositeAction) Execute() {
	// run the composited actions
	for _, action := range c.actions {
		action.Execute()
	}
}

// AddAction adds an action to the actions collection
func (c *CompositeAction) AddAction(action Action) {
	c.actions = append(c.actions, action)
}
