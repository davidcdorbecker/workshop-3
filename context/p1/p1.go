package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	//background()
	//withCancel()
	//withTimeOut()
	withValue()
}

func background() {
	ctx := context.Background()

	sleepAndTalk(ctx, 3*time.Second, "Hello background!")
}

func withCancel() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	sleepAndTalk(ctx, 3*time.Second, "Hello withCancel!")
}

func withTimeOut() {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))
	defer cancel()

	sleepAndTalk(ctx, 3*time.Second, "Hello withTimeOut!")
}

func withValue() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ctx-key", " with value")
	sleepAndTalk(ctx, time.Second, "hello")
	//
}

func sleepAndTalk(ctx context.Context, d time.Duration, s string) {
	msg := ctx.Value("ctx-key").(string)
	select {
	case <-time.After(d):
		fmt.Println(s, msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
