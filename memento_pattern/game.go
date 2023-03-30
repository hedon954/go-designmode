package main

import (
	"fmt"
)

type Player struct {
	x, y int // position
	hp   int // health point
}

// Memento is the memento object, use to hold the player status
type Memento struct {
	Player
}

// Originator is the originator object, use to save or restore the player status
type Originator struct {
	Player
}

func (o *Originator) MoveTo(x, y int) {
	o.x, o.y = x, y
	fmt.Printf("player moves to (%d, %d)\n", x, y)
}

func (o *Originator) TakeDamage(damage int) {
	o.hp -= damage
	fmt.Printf("player takes %d damage, now the hp is %d\n", damage, o.hp)
}

// CreateMemento creates a memento according to current player's status
func (o *Originator) CreateMemento() *Memento {
	return &Memento{
		Player{
			x:  o.x,
			y:  o.y,
			hp: o.hp,
		},
	}
}

// RestoreMemento restores player's status from memento
func (o *Originator) RestoreMemento(m *Memento) {
	o.x, o.y, o.hp = m.x, m.y, m.hp
}

// CareTaker is the manager object
type CareTaker struct {
	mementos []*Memento
}

func (ct *CareTaker) AddMementos(ms ...*Memento) {
	ct.mementos = append(ct.mementos, ms...)
}

func (ct *CareTaker) GetLastMemento() *Memento {
	n := len(ct.mementos)
	if n == 0 {
		return nil
	}
	return ct.mementos[n-1]
}

func main() {
	// create a player
	p := Player{0, 0, 100}

	// create originator and manager
	originator := Originator{
		Player: p,
	}
	careTaker := &CareTaker{}

	// player move and save status
	originator.MoveTo(1, 1)
	careTaker.AddMementos(originator.CreateMemento())

	// player is damaged and save status
	originator.TakeDamage(20)

	// restore status
	m := careTaker.GetLastMemento()
	originator.RestoreMemento(m)
	fmt.Printf("player restores to (%d,%d), and hp is %d\n", originator.x, originator.y, originator.hp)
}
