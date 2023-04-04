package main

import (
	"context"
	"fmt"
	"time"
)

// We can use the Context to everything that do I/O in go. HTTP calls, TCP calls and so on...
func main() {
	// ctx := context.TODO() => EXPLICITY SAYS, CTX DO NOTHING BUT WE WILL IN THE FUTURE IMPLEMENTS IT
	// tree like data structure
	// white ctx, root parent ctx
	ctx := context.Background()
	// cancel function, when execte cancel() this will propagates to every ctx child...

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		// if less then 5 seconds cancel will be trigered and case "<-ctx.Done()" executed
		// simulates some I/O
		time.Sleep(time.Second * 10)
		cancel()
	}()

	bookHotel(ctx)
}

// good practices, if uses ctx always pass as first arg
func bookHotel(ctx context.Context) {
	// waits for some channel
	select {
	// when its done
	case <-ctx.Done():
		fmt.Println("Exceeded time to book the room")
	case <-time.After(time.Second * 5):
		fmt.Println("Room reserved")
	}
}
