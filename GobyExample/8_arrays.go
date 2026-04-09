package main // idk why we right this asofnow

import "fmt"

func main () {
	var a [5]int
	fmt.Println("emp:", a) // emp: [0 0 0 0 0]
	
	// set value at index
	a[4] = 100
	fmt.Println("set:", a) // set: [0 0 0 0 100]
	
	// array[index] = value syntax
	fmt.Println("get:", a[4])  // get: 100
	
	fmt.Println("len:", len(a)) // 5
	
	// declare and intialize in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b) // dcl [1 2 3 4 5]
	
	// using ... complier counts the number of elements.
	b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
    
    // If you specify the index with :, the elements in between will be zeroed.
    b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b) // idx: [100 0 0 400 500]
    
    // Array types are one-dimensional, but you can compose types
    // to build multi-dimensional data structures.
    var twoD [2][3]int
    for i := range 2 {
    	for j:= range 3 {
     		twoD[i][j] = i + j;
     	}
    }
    fmt.Println("2d: ", twoD) // 2d:  [[0 1 2] [1 2 3]]
    
    // You can create and initialize
    // multi-dimensional arrays at once too.
    
    twoD = [2][3]int{
    {1, 2, 3},
    {1, 2, 3},
    }
    fmt.Println("2d: ", twoD) // 2d:  [[1 2 3] [1 2 3]]
}