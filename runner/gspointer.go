package runner

import "fmt"

func runPointer() {
	printStart()
	defer printEnd()
	myString := "Hello world"
	fmt.Printf("myString @ %x, v = %s, v[0] = %c\n", myString, myString, myString[0])
	myString = changeString(myString)
	fmt.Printf("myString @ %x, v = %s, v[0] = %c\n", myString, myString, myString[0])
	changeString2(&myString)
	fmt.Printf("myString @ %x, v = %s, v[0] = %c\n", myString, myString, myString[0])

	myInt := int64(1)
	fmt.Printf("myInt @ %x, v = %d\n", &myInt, myInt)
	myInt = changeInt(myInt)
	changeInt2(&myInt)
}

func changeInt(input int64) int64 {
	input = 11
	fmt.Printf("input @ %x, v = %d\n", &input, input)
	return input
}

func changeInt2(input *int64) {
	fmt.Printf("input @ %x, v = %d\n", input, *input)
	*input = 111
}

func changeString(input string) string {
	fmt.Printf("input @ %x, v = %s\n", input, input)
	input = "Hello world !"

	return input
}

func changeString2(input *string) {
	fmt.Printf("input @ %x, v = %s\n", input, *input)
	*input = "Hello world !!"
}
