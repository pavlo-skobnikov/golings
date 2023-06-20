// structs2
// Make me compile!
package main

import "fmt"

type Person struct {
	// don't just create the phone field here. embed a new struct
	name string
	age  int
}

type PersonWithPhoneNumber struct {
  Person
  phone string
}

func main() {
  person := PersonWithPhoneNumber{
    Person: Person{name: "John", age: 32},
    phone: "123-456-7890",
  }

	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
