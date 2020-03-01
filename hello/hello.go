package main

import (
	"fmt"

	"math"

	"github.com/Allah-The-Dev/stringutil"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(stringutil.Reverse("hseleeN"))
}
