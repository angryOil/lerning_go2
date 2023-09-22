package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var jobList [10]job
	for i := 0; i < 10; i++ {
		jobList[i] = &squareJob{i}
	}
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		nowJob := jobList[i]
		go func() {
			nowJob.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}

type job interface {
	Do()
}

type squareJob struct {
	index int
}

func (j *squareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과 %d \n", j.index, j.index*j.index)
}
