package runner

import (
	"container/list"
	"fmt"
)

func runArray() {
	printStart()
	var a [10]uint64
	a[0] = 1
	b := []int{1, 2, 3}
	fmt.Printf("Array Cap %d\n", cap(a))
	fmt.Printf("Array Len %d\n", len(b))
	fmt.Println("-------------------")

	l := list.New()
	e4 := l.PushBack(MyData{4, "4"})
	e1 := l.PushFront(MyData{1, "1"})
	l.InsertBefore(MyData{3, "3"}, e4)
	l.InsertAfter(MyData{2, "2"}, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value.(MyData).Id)
	}
	fmt.Printf("\n")
	printEnd()
}
