package main

import (
	"context"
	"fmt"
)

type subscriber struct {
	ctx   context.Context
	name  string
	msgCh chan string
}

func newSubscriber(name string, ctx context.Context) *subscriber {
	return &subscriber{
		ctx:   ctx,
		name:  name,
		msgCh: make(chan string),
	}
}

func (s *subscriber) subscribe(pub *publisher) {
	pub.subscribe(s.msgCh)
}

func (s *subscriber) update() {
	for {
		select {
		case msg := <-s.msgCh:
			fmt.Println(s.name, " got msg:", msg)
		case <-s.ctx.Done():
			wg.Done()
			return
		}
	}
}
