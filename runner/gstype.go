package runner

import (
	"fmt"
	"unsafe"
)

type MyData struct {
	Id     uint8
	Field1 string
}

func (md *MyData) greet() string {
	return "Hi, My id is " + md.Field1
}

func runType() {
	printStart()
	const someText = "string"
	a := MyData{Id: 2, Field1: someText}
	fmt.Printf("%T size = %d\n", a.Id, unsafe.Sizeof(a.Id))
	fmt.Printf("%T size = %d\n", a.Field1, unsafe.Sizeof(a.Field1))
	a.Id = 1
	fmt.Printf("greet = %s\n", a.greet())
	printEnd()
}
