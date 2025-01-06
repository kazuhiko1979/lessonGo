// Go:build main

package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("deferred call in foo")

	fmt.Println("hello from foo")
}

func main() {

	// defer fmt.Println("deferred call in main")
	// foo()

	// fmt.Println("hello")

	// fmt.Println("run")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("success")

	file, _ := os.Open("lesson.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))

}
