// Explicitly Declared and used by the compiler : to check type-correctness of function call
package main
import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)
	
	// var - declares 1 or more variables
    var b, c int = 1, 2 // can declare multiple variables at once.
	fmt.Println(b, c)
	
	
	var d = true
	fmt.Println(d)
	
	// declare but not initialized -> 0
	var e int
	fmt.Println(e)
	
	// shorthand for declaring and initializing var
	f := "apple"
	fmt.Println(f)
}