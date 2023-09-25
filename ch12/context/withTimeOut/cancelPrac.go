package main

import "context"

func longRunningThingManger(ctx context.Context, data string) (string, error) {
	type wrapper struct {
		result string
		err    error
	}

	ch := make(chan wrapper, 1)
	go func() {
		ch <- wrapper{
			result: "",
			err:    nil,
		}
	}()
	select {
	case data := <-ch:
		return data.result, data.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
