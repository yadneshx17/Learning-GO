package main

import "fmt"

// algebric data type -> composite datatype -> type form by combining other types.
// Enumerated Types -> is a type that has fixed number of possible values, each with distinct value.
// Go doesnt have an enum type implicity as a language feature,

// iota a concept or a predefined variable use with const declaration to generate a sequence of related, increasing integer constants (0, 1, 2, 3)

// custom type to prevent accidentally passing a raw integer where your enum was expected.
// line: 33 in the switch case function
type ServerState int

const (
	StateIdle ServerState = iota // creates successive const values. 0, 1, 2
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func main() {
	ns := transition(StateConnected)
	fmt.Println(ns)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
