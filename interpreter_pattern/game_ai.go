package main

import (
	"fmt"
	"math/rand"
)

// Node is the node of decision-making tree
type Node interface {

	// Evaluete is used to calculate the value of current node
	Evaluate() bool
}

// ConditionNode is a condition node
type ConditionNode struct {
	Condition func() bool
}

func (node *ConditionNode) Evaluate() bool {
	return node.Condition()
}

// ActionNode is an action node
type ActionNode struct {
	Action func()
}

func (node *ActionNode) Evaluate() bool {
	node.Action()
	return true
}

// CompositeNode is a composie node who contains a lot of nodes
type CompositeNode struct {
	Children []Node
}

func (node *CompositeNode) Evaluate() bool {
	for _, n := range node.Children {
		if !n.Evaluate() {
			return false
		}
	}
	return true
}

// buildDecisionTree builds a decision tree by using interpreter pattern
func buildDecisionTree() Node {
	return &CompositeNode{
		Children: []Node{
			&ConditionNode{
				Condition: func() bool {
					// Randomly decide whether to attack
					return rand.Intn(2) == 0
				},
			},
			&ActionNode{
				Action: func() {
					// attack enemy
					fmt.Println("attack enemy")
				},
			},
		},
	}
}
