package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var client = http.Client{}

func callBoth(ctx context.Context, errVal string, slowUrl string, fastUrl string) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := callServer(ctx, "slow", slowUrl)
		if err != nil {
			cancel()
		}
	}()

	go func() {
		defer wg.Done()
		err := callServer(ctx, "fast", fastUrl+"?error="+errVal)
		if err != nil {
			cancel()
		}
	}()
	wg.Wait()
	fmt.Println("done with both")
}

func callServer(ctx context.Context, label string, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(label, "request err", err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(label, "res err", err)
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(label, "read err:", err)
		return err
	}
	result := string(data)
	if result != "" {
		fmt.Println(label, "result:", result)
	}
	if result == "error" {
		fmt.Println("cancelling from label", result)
		return errors.New("error happened")
	}
	return nil
}
