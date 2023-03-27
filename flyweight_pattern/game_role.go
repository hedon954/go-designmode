package main

import (
	"fmt"
	"sync"
)

// Role is the abstract flyweight interface, it defines the common properties and methods
type Role interface {
	GetName() string
	GetSpeed() int
	GetPower() int
	SetSpeed(speed int)
	SetPower(power int)
	UseSkill()
}

// RoleImpl is the concrete flyweight instance
type RoleImpl struct {
	name  string // external status
	speed int    // internal status
	power int    // internal status
}

func (r *RoleImpl) GetName() string {
	return r.name
}

func (r *RoleImpl) GetSpeed() int {
	return r.speed
}

func (r *RoleImpl) GetPower() int {
	return r.power
}

func (r *RoleImpl) SetSpeed(speed int) {
	r.speed = speed
}

func (r *RoleImpl) SetPower(power int) {
	r.power = power
}

func (r *RoleImpl) UseSkill() {
	fmt.Printf("%s uses skill\n", r.name)
}

// RoleFactory is the flyweight factory, it holds all the flyweight objects in memory
type RoleFactory struct {
	roles sync.Map
}

func (rf *RoleFactory) GetRole(name string) Role {

	// if exists, return it
	if role, ok := rf.roles.Load(name); ok {
		return role.(Role)
	}

	// first time to invoke, create it
	var newRole Role
	switch name {
	case "warrior":
		newRole = &RoleImpl{
			name:  "warrior",
			speed: 5,
			power: 10,
		}
	case "mage":
		newRole = &RoleImpl{
			name:  "mage",
			speed: 3,
			power: 15,
		}
	case "archer":
		newRole = &RoleImpl{
			name:  "archer",
			speed: 8,
			power: 8,
		}
	}

	// use loadOrStore to determine if any other goroutine has been created the role
	actual, loaded := rf.roles.LoadOrStore(name, newRole)
	if loaded {
		return actual.(Role)
	}
	return newRole
}
