package main

import (
	"fmt"
	"strings"
)

type Handler interface {
	HandleRequest(req string)
}

func GetHandler() Handler {
	lastHandler := NewLastHandler(nil)
	handlerB := NewConcreteHandlerB(lastHandler)
	handlerA := NewConcreteHandlerA(handlerB)
	return handlerA
}

type ConcreteHandlerA struct {
	successor Handler
}

func NewConcreteHandlerA(h Handler) *ConcreteHandlerA {
	return &ConcreteHandlerA{
		successor: h,
	}
}

func (cha *ConcreteHandlerA) HandleRequest(req string) {
	if strings.Contains(req, "A") {
		fmt.Println("ConcreteHandlerA handle request")
		return
	}

	if cha.successor != nil {
		cha.successor.HandleRequest(req)
	}
}

type ConcreteHandlerB struct {
	successor Handler
}

func NewConcreteHandlerB(h Handler) *ConcreteHandlerB {
	return &ConcreteHandlerB{
		successor: h,
	}
}

func (chb *ConcreteHandlerB) HandleRequest(req string) {
	if strings.Contains(req, "B") {
		fmt.Println("ConcreteHandlerB handle request")
		return
	}

	if chb.successor != nil {
		chb.successor.HandleRequest(req)
	}
}

type LastHandler struct {
	successor Handler
}

func NewLastHandler(h Handler) *LastHandler {
	return &LastHandler{
		successor: h,
	}
}

func (lh *LastHandler) HandleRequest(req string) {
	fmt.Println("LastHandler handle request")
}
