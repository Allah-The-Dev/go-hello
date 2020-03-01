package array

import "fmt"

// HandleArray ... method to showcase array in go and related methods
func HandleArray() {
	// declaring the array
	var numArray [3]int

	// initializing array with values
	numArray[0] = 0
	numArray[1] = 1
	numArray[2] = 2

	// printing some values
	fmt.Println(numArray[2])
	fmt.Println(len(numArray))
	fmt.Println(numArray)

	// array dec and initializing without mentioning size
	directions := [...]int{1, 2, 3, 4, 5}
	fmt.Println(directions)
}
