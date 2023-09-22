package main

import "context"

// chan a 양방향 채널
// chan <- a 쓰기 전용(write only
// <- chan a 읽기 전용 read only

type publisher struct {
	ctx         context.Context
	subscribeCh chan chan<- string
	publishCh   chan string
	subscribers []chan<- string
}

func newPublisher(ctx context.Context) *publisher {
	return &publisher{
		ctx:         ctx,
		subscribeCh: make(chan chan<- string),
		publishCh:   make(chan string),
		subscribers: make([]chan<- string, 0),
	}
}

func (p *publisher) subscribe(sub chan<- string) {
	p.subscribeCh <- sub
}

func (p *publisher) publish(msg string) {
	p.publishCh <- msg
}

func (p *publisher) update() {
	for {
		select {
		case sub := <-p.subscribeCh:
			p.subscribers = append(p.subscribers, sub)
		case msg := <-p.publishCh:
			for _, subscriber := range p.subscribers {
				subscriber <- msg
			}
		case <-p.ctx.Done():
			wg.Done()
			return
		}
	}
}
