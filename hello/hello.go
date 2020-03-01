package main

import (
	"fmt"

	"math"

	"github.com/Allah-The-Dev/stringutil"

	"github.com/Allah-The-Dev/array"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {

	//using math functions
	fmt.Println(sqrt(2), sqrt(-4))

	//using string
	fmt.Println(stringutil.Reverse("hseleeN"))

	// array in go
	array.HandleArray()
}
