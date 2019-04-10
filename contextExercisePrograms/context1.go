package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctxBackground := context.Background()

	/*ctx, cancel := context.WithTimeout(ctxBackground, time.Second)
	defer cancel() //(send deadline to context), log.Print(ctx.Err()) > 2019/02/04 12:27:57 context deadline exceeded
	//cancel() //(send cancel signal to context), log.Print(ctx.Err()) > 2019/02/04 12:28:25 context canceled
	*/
	//OR
	/*ctx, cancel := context.WithCancel(ctxBackground)
	time.AfterFunc(time.Second, cancel) //(send cancel signal to context) 2019/02/04 12:33:44 context canceled
	*/
	//OR
	ctx, cancel := context.WithCancel(ctxBackground)
	go func() {
		time.Sleep(time.Second)
		cancel() //(send cancel signal to context), log.Print(ctx.Err()) > 2019/02/04 12:32:54 context canceled
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
