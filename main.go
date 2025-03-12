package main

import (
	"fmt"
	"log"
	"time"
)

type startGoroutineFn func(done <-chan interface{}, pulseInterval time.Duration) (heartbeat <-chan interface{})

func DoWorkFn(done <-chan interface{}, intList ...int) (startGoroutineFn, <-chan interface{}) {

	intStream := make(chan interface{})

	doWork := func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} {
		heartbeat := make(chan interface{})

		go func() {
			defer close(heartbeat)
			defer close(intStream)

			pulse := time.Tick(pulseInterval)

		valueLoop:
			for _, i := range intList {
				if i < 0 {
					log.Printf("Warning: negative value detected: %v. Skipping...\n", i)
					// return
					continue
				}
				for {
					select {
					case <-pulse:
						select {
						case heartbeat <- struct{}{}:
						default:
						}
					case intStream <- i:
						continue valueLoop
					case <-done:
						return
					}
				}
			}
		}()
		return heartbeat
	}
	return doWork, intStream
}

func newSteward(timeout time.Duration, maxRestarts int, startGoroutineFn startGoroutineFn) startGoroutineFn {
	return func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} {
		heartbeat := make(chan interface{})
		go func() {
			defer close(heartbeat)
			var wardDone chan interface{}
			var wardHeartbeat <-chan interface{}
			startWard := func() {
				wardDone = make(chan interface{})
				wardHeartbeat = startGoroutineFn(or(wardDone, done), timeout/2)
			}

			startWard()
			pulse := time.Tick(pulseInterval)

			restartCount := 0

		monitorLoop:
			for {
				if restartCount >= maxRestarts {
					log.Println("steward: max restarts reached, stopping monitoring.")
					close(wardDone)
					return
				}

				timeoutSignal := time.After(timeout)
				for {
					select {
					case <-pulse:
						select {
						case heartbeat <- struct{}{}:
						default:
						}
					case <-wardHeartbeat:
						continue monitorLoop
					case <-timeoutSignal:
						log.Println("steward: ward unhealthy; restarting")
						close(wardDone)
						restartCount++
						startWard()
						continue monitorLoop
					case <-done:
						return
					}
				}
			}
		}()
		return heartbeat
	}
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

func main() {

	done := make(chan interface{})
	defer close(done)

	doWork, intStream := DoWorkFn(done, 1, 2, -1, -2, -4, -5, 3, 4, 5)

	doWorkWithSteward := newSteward(1*time.Second, 3, doWork)
	doWorkWithSteward(done, 1*time.Second)

	for i := range intStream {
		fmt.Printf("Received: %v\n", i)
	}

	fmt.Println("Processing completed. Exiting...")

}
