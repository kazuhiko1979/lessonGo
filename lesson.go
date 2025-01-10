// Go:build main
package main

import "fmt"

type Vertex struct {
	X, Y int
}

// func Plus(v Vertex) int {
// 	return v.X + v.Y
// }

// func (v Vertex) Plus() int {
// 	return v.X + v.Y
// }

func (v Vertex) String() string {
	return fmt.Sprintf("Q2 X is %d! Y: %d!", v.X, v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
}
