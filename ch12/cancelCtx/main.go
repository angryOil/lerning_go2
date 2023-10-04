package main

import (
	"context"
	"fmt"
)

func main() {
	ss := slowServer()
	defer ss.Close()
	fs := fastServer()
	defer fs.Close()

	ctx := context.Background()
	fmt.Println(ss.URL, fs.URL)
	callBoth(ctx, "true", ss.URL, fs.URL)
}
