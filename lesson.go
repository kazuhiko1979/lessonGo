package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int) // 15, 15 queue
	go goroutine1(s, c)
	go goroutine1(s, c)
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)

}

// // Go:build main
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func goroutine(s string, wg *sync.WaitGroup) {
// 	for i := 0; i < 5; i++ {
// 		// time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// 	wg.Done()
// }

// func normal(s string) {
// 	for i := 0; i < 5; i++ {
// 		// time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}

// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	wg.Add(1)

// 	go goroutine("world", &wg)
// 	normal("hello")
// 	// time.Sleep(2000 * time.Millisecond)
// 	wg.Wait()

// }
