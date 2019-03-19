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
	ctxBackground := context.Background()
	ctx, cancel := context.WithCancel(ctxBackground)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()
	sleepAndTalk(ctx, 5*time.Second, "hello")
}

func sleepAndTalk(ctx context.Context, duration time.Duration, msg string) {
	select {
	case <-time.After(duration):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}
