package main

type Handler interface {
	Execute(*patient) error
	SetNext(Handler) Handler
	Do(*patient) error
}

type Next struct {
	nextHandler Handler
}

func (n *Next) SetNext(handler Handler) Handler {
	n.nextHandler = handler
	return handler
}

func (n *Next) Execute(patient *patient) error {
	if n.nextHandler != nil {
		if err := n.nextHandler.Do(patient); err != nil {
			return err
		}
		return n.nextHandler.Execute(patient)
	}
	return nil
}
