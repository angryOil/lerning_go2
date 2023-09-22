package main

import "fmt"

func main() {
	m := &mail{}
	var listener eventListener = &alarm{}

	m.Register(listener)
	m.OnRecv()
}

type eventListener interface {
	OnFire()
}

type alarm struct {
}

func (a alarm) OnFire() {
	fmt.Println("메일이 왔습니다")
}

type event interface {
	Register(eventListener)
}

type mail struct {
	listener eventListener
}

func (m *mail) Register(listener eventListener) {
	m.listener = listener
}

func (m *mail) OnRecv() {
	m.listener.OnFire()
}
