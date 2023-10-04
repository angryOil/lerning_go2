package withTimeOut

import (
	"context"
	"time"
)

func doPrentCtxCancelAndChildCtxCancel() string {
	ctx := context.Background()
	parent, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	child, cancel2 := context.WithTimeout(parent, 3*time.Second)
	defer cancel2()
	start := time.Now()
	<-child.Done()
	end := time.Now()
	return end.Sub(start).String()
}
