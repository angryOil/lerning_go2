package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println(nil)
	}
	req.Header.Add("X-My-Client", "learning go")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic("unexpected status:got" + res.Status)
	}
	fmt.Println(res.Header.Get("Content-Type"))

	//bodyByte, err := io.ReadAll(res.Body)

	var data struct {
		UserId    int    `json:"userId"`
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	json.NewDecoder(res.Body).Decode(&data)
	fmt.Println(data)
	fmt.Printf("%+v\n", data)
}
