package runner

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func runError() {
	printStart()
	defer printEnd()
	resultString, _ := greetUser("asd")
	fmt.Println(resultString)
	_, resultError2 := greetUser("")
	fmt.Println(resultError2)
	_, resultError3 := greetUser("abc")
	fmt.Println(resultError3)
	maybeWePanic()
}

type ourError struct {
	name string
}

func (e ourError) Error() string {
	return e.name
}

func maybeWePanic() {
	rand.Seed(time.Now().UnixNano())
	v := rand.Int()
	if v%7 == 0 {
		panic("what just happened ?")
	}
}

func greetUser(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	if name == "abc" {
		return "", ourError{name: "our error"}
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
