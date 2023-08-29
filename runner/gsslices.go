package runner

import (
	"fmt"
)

func runSlice() {
	printStart()
	a := [3]string{
		"one",
		"two",
		"three",
	}
	b := make([]int, 1)
	b[0] = 1
	b = append(b, 2)
	fmt.Printf("slice cap = %d, slice len = %d\n", cap(b), len(b))

	for _, v := range a[0:3] {
		fmt.Println(v)
	}

	for _, v := range a[1:2] {
		fmt.Println(v)
	}
	printEnd()
}
