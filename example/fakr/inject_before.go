package fakr

import "fmt"

func boo() {
	fmt.Println("hello")
}

// inject
func bar() {
	fmt.Println("hello")
}
