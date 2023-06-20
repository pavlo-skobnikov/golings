// generics1
// Make me compile!

package main

import "fmt"

func main() {
	print("Hello, World!")
	print(42)
}

func print[T interface{}](value T) {
	fmt.Println(value)
}
