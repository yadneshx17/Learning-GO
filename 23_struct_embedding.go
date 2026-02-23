package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "porsche",
	}

	// can access base's field directly on co
	fmt.Printf("co={num: %v and str: %v}\n", co.num, co.str)

	// can also use the full path. using embedded type name
	fmt.Println("also num:", co.base.num)

	// Since container embeds base, the methods of base also become methods of a container.
	// invokes a method that was embedded from `base` and direclty on `co`.
	fmt.Println("describe:", co.describe())

	// i didn't understood this.
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
