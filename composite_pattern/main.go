package main

import (
	"time"
)

func main() {
	action1 := &CompositeAction{}
	action11 := &PlayerAction{
		playerId:   11,
		actionTime: time.Now(),
	}
	action12 := &PlayerAction{
		playerId:   12,
		actionTime: time.Now(),
	}
	action1.AddAction(action11)
	action1.AddAction(action12)

	action2 := &CompositeAction{}
	action21 := &PlayerAction{
		playerId:   21,
		actionTime: time.Now(),
	}
	action22 := &PlayerAction{
		playerId:   22,
		actionTime: time.Now(),
	}
	action2.AddAction(action21)
	action2.AddAction(action22)

	composite := &CompositeAction{}
	composite.AddAction(action1)
	composite.AddAction(action2)

	composite.Execute()
}
