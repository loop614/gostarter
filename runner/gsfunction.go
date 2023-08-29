package runner

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func runFunction() {
	printStart()
	defer printEnd()
	myString := concat("Hello ", "World !")
	fmt.Println(myString)
	myString2 := concatMany("Hello ", "World", "!")
	fmt.Println(myString2)
	myString3 := callMe(concat, "Hello ", "World !")
	fmt.Println(myString3)
	naturalNumbers := map[string]int64{
		"first":  34,
		"second": 12,
	}
	realNumbers := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	myResult1 := sumNumbers[string, int64](naturalNumbers)
	fmt.Println(myResult1)

	myResult2 := sumNumbers[string, float64](realNumbers)
	fmt.Println(myResult2)

	myResult3 := min(1, 2)
	fmt.Println(myResult3)

	myResult4 := min(1.1, 2.2)
	fmt.Println(myResult4)

	myResult5 := min("abc", "cba")
	fmt.Println(myResult5)
}

func sumNumbers[K comparable, V int64 | float64](numbers map[K]V) V {
	var sum V
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func min[T constraints.Ordered](input1, input2 T) T {
	if input1 < input2 {
		return input1
	}
	return input2
}

func concat(s1, s2 string) string {
	return s1 + s2
}

func concatMany(ss ...string) string {
	res := ""
	for i := 0; i < len(ss); i++ {
		res += ss[i]
	}
	return res
}

func callMe(
	concat func(s1 string, s2 string) string,
	input1 string,
	input2 string,
) string {
	return concat(input1, input2)
}
