package main

import (
	"fmt"
)

type State interface {
	Update(role *Role)
}

type Role struct {
	state State
}

// SetState sets the role's state
func (r *Role) SetState(state State) {
	r.state = state
}

// Update updates the role's state
func (r *Role) Update() {
	r.state.Update(r)
}

type NormalState struct{}

func (s *NormalState) Update(role *Role) {
	fmt.Println("role is in normal state")
}

// 受伤状态
type InjuredState struct{}

func (s *InjuredState) Update(role *Role) {
	// 受伤状态下的行为
	fmt.Println("role is in injured state")
}

type DeadState struct{}

func (s *DeadState) Update(role *Role) {
	// 死亡状态下的行为
	fmt.Println("role is in dead state")
}

func main() {
	role := &Role{}
	normalState := &NormalState{}
	injuredState := &InjuredState{}
	deadState := &DeadState{}

	role.SetState(normalState)
	role.Update()

	role.SetState(injuredState)
	role.Update()

	role.SetState(deadState)
	role.Update()
}