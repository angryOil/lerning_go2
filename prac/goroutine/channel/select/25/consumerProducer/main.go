package main

import (
	"fmt"
	"sync"
	"time"
)

// ch <- ch <- ch

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	tireCh := make(chan *car)
	paintCh := make(chan *car)

	fmt.Println("start factory")
	wg.Add(3)
	go makeBody(tireCh)
	go installTire(tireCh, paintCh)
	go paintCar(paintCh)

	wg.Wait()
	fmt.Println("close the factory")
}

func paintCar(paintCh chan *car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "red"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f complete car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	car := <-paintCh
	fmt.Println(car)
	wg.Done()
}

func installTire(tireCh chan *car, paintCh chan *car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "summer tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func makeBody(tireCh chan *car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			car := &car{}
			car.Body = "sport car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

type car struct {
	Body  string
	Tire  string
	Color string
}
