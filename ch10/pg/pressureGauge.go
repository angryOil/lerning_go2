package main

import (
	"errors"
	"net/http"
)

func main() {
	pg := newPg(3)
	http.HandleFunc("/req", func(w http.ResponseWriter, r *http.Request) {
		err := pg.process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("too many request"))
		}
	})
	http.ListenAndServe(":8080", nil)
}

func doThingThatShouldBeLimited() string {
	//time.Sleep(1 * time.Second)
	return "돈"
}

type pressureGauge struct {
	ch chan struct{}
}

func newPg(limit int) *pressureGauge {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &pressureGauge{
		ch: ch,
	}
}

func (pg *pressureGauge) process(f func()) error {
	select {
	case <-pg.ch:
		f()
		//pg.ch <- struct{}{}
		return nil
	default:
		return errors.New("no more cap")
	}
}
