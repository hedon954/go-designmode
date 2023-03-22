package main

import (
	"fmt"
)

// patient defines the process of seeing a patient
type patient struct {
	Name              string
	ReceptionDone     bool
	DockerCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}

type Start struct {
	Next
}

func (s *Start) Do(p *patient) error {
	// Start dose nothing, just used as the first handler to transfer the request to Next Handler
	return nil
}

type Reception struct {
	Next
}

func (r *Reception) Do(p *patient) error {
	if p.ReceptionDone {
		return nil
	}
	fmt.Println("Reception...")
	p.ReceptionDone = true
	return nil
}

type DockerCheck struct {
	Next
}

func (d *DockerCheck) Do(p *patient) error {
	if p.DockerCheckUpDone {
		return nil
	}
	fmt.Println("docker check...")
	p.DockerCheckUpDone = true
	return nil
}

type Payment struct {
	Next
}

func (p *Payment) Do(p2 *patient) error {
	if p2.PaymentDone {
		return nil
	}
	fmt.Println("payment...")
	p2.PaymentDone = true
	return nil
}

type Medicine struct {
	Next
}

func (m Medicine) Do(p *patient) error {
	if p.MedicineDone {
		return nil
	}
	fmt.Println("medicine...")
	p.MedicineDone = true
	return nil
}
