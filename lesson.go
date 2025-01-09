// Go:build main
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %v.", p.Name)
}

func main() {
	mike := Person{Name: "Mike", Age: 22}
	fmt.Println(mike)

}
