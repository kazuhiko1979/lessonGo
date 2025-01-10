// Go:build main
package main

import "fmt"

type UserNotFound struct {
	username string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.username)
}

func myFunc() error {
	// Something wrong
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{username: "mike"}

}

func main() {
	e1 := &UserNotFound{"mike"}
	e2 := &UserNotFound{"mike"}
	fmt.Println(e1 == e2)

	err := myFunc()
	if err != nil {
		fmt.Println(err)
	}
}
