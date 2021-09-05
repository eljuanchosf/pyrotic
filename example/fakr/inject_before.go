package fakr

import "fmt"

func booAfter() {
	fmt.Println("hello")
}

// inject
func barAfter() {
	fmt.Println("hello")
}
