package main

import "fmt"

func main() {

	type Person struct {
		Name  string
		Likes []string
	}

	people := []*Person{
		{Name: "Alice", Likes: []string{"Go", "Music"}},
		{Name: "Bob", Likes: []string{"Go", "Sports"}},
	}

	likes := make(map[string][]*Person)

	for _, p := range people {
		for _, l := range p.Likes {
			// If likes[l] doesn’t exist → Go gives nil
			// append(nil, p) works → creates a new slice automatically
			// []*Person{p}
			// likes[l] = []*Person{p}
			likes[l] = append(likes[l], p)
			for _, person := range likes[l] {
				if person.Name == p.Name {
					fmt.Println(p.Name, "likes", l)
				}
			}
		}
	}
}
