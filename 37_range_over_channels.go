package main 

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	
	// can close the nonempty channel but still receive remaining values.
	close(queue)
	
	for elem:= range queue {
		fmt.Println(elem)
	}
}