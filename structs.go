package main

import "fmt"

// struct - typed collection of fields
// go is a garbage collected language, only be cleanup when there is no active references pointing to it.

type person struct {
	name string
	age  int
}

// constructs a new person struct with the given name
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 19
	return &p // pointer is returned of local variable
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})          // omitted field will have zero value.
	fmt.Println(&person{name: "Ann", age: 40}) // `&` yeilds pointer to the struct.
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	// If a struct type is only used for a single value,
	// we don’t have to give it a name.
	// The value can have an anonymous struct type.
	dog := struct {
		name string
		age  bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
