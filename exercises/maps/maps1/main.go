// maps1
// Make me compile!
//
package main

import "fmt"

func main() {
	// Map with people names and their ages
	m := make(map[string]int)

  const john = "John"
  const ana = "Ana"

	m[john] = 30
	m[ana] = 21

	fmt.Printf("John is %d and Ana is %d", m[john], m[ana])
}
