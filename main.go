package main

import (
	"fmt"
	"time"
)

func DoWorkVer1(done <-chan interface{}, palseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
	heartbeat := make(chan interface{})
	results := make(chan time.Time)

	go func() {
		// defer close(heartbeat)
		// defer close(results)

		heartbeatpalse := time.Tick(palseInterval)
		workGen := time.Tick(2 * palseInterval)

		sendHeartBeatPalse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
		}

		sendResult := func(result time.Time) {
			for {
				select {
				case <-done:
					return
				case <-heartbeatpalse:
					sendHeartBeatPalse()
				case results <- result.Local():
					return
				}
			}
		}

		// for {
		for i := 0; i < 2; i++ {
			select {
			case <-done:
				return
			case <-heartbeatpalse:
				sendHeartBeatPalse()

			case result := <-workGen:
				sendResult(result)
			}
		}
	}()

	return heartbeat, results
}

func main() {

	done := make(chan interface{})

	const timeout = 2 * time.Second
	heartbeat, results := DoWorkVer1(done, timeout/2)

	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				return
			}
			fmt.Println("receive heartbeat")

		case r, ok := <-results:
			if !ok {
				return
			}
			fmt.Printf("receive result %v\n", r)

		case <-time.After(timeout):
			fmt.Println("worker gorutine is not healthy!")
			return
		}
	}
}
