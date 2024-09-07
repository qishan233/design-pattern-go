package main

import (
	"fmt"

	"math/rand"
)

type Receiver interface {
	Action() string
}

type ConcreteReceiver struct{}

func (cra *ConcreteReceiver) Action() string {
	return fmt.Sprintf("ConcreteReceiverA %d", rand.Intn(100))
}
