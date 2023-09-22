package main

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	useHttp()
}

func useHttp() {
	pg := newP(10)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.process(func() {
			fmt.Println("hllo")
			fmt.Println(len(pg.ch))
			w.Write([]byte(doThingThatShouldBeLimited()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("too many requests"))
		}
	})
	http.ListenAndServe(":8080", nil)
}

func doThingThatShouldBeLimited() string {
	time.Sleep(2 * time.Second)
	return "done"
}

func useWg() {
	p := newP(1)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			err := p.process(func() {
				fmt.Println("id:", id, " hello?")
			})
			if err != nil {
				fmt.Println("id:", id, " is err", err)
			}
		}(i)
	}
	wg.Wait()
}

type pressureGauge struct {
	ch chan struct{}
}

func newP(limit int) *pressureGauge {
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
		pg.ch <- struct{}{}
		return nil
	default:
		return errors.New("no more capacity")
	}
}
