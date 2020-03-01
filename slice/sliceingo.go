package slice

import "fmt"

// HandleSliceBasic ... declaration and methods on slice
func HandleSliceBasic() {
	// array
	a := [5]string{"one", "two", "three", "four", "five"}
	fmt.Println("Array after creation : ", a)

	//slice
	var b []string = a[1:4] //[inclusive:exclusinve]
	fmt.Println("Slice after creation : ", b)

	//changing slice data
	b[0] = "changed"
	fmt.Println("Array after change in slice : ", a)
	fmt.Println("Slice after change in slice element : ", b)
}

// MethodsOnSlice ... len and append methods on slice
func MethodsOnSlice() {
	// array
	a := [5]string{"one", "two", "three", "four", "five"}
	sliceA := a[1:3]
	b := [...]string{"1", "2", "3", "4", "5"}
	sliceB := b[2:5]

	fmt.Println("Lenght of silce_a : ", len(sliceA))
	fmt.Println("Lenght of silce_b : ", len(sliceB))

	sliceA = append(sliceA, sliceB...)
	fmt.Println("sice_a after appending slice_b : ", sliceA)
	fmt.Println("array a after change in slice_a : ", a)

	sliceB = append(sliceB, "6")
	fmt.Println("sice_b after appending 6 : ", sliceB)
	fmt.Println("array b after change in slice_b : ", b)
}
