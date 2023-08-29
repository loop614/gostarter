package runner

import "fmt"

func runMap() {
	printStart()
	fmt.Println("Maps are great")
	defer printEnd()
	userOneId := userId{id: 1, specialNumber: 2}
	users := map[userId]user{
		userOneId: {id: 1, name: "name", address: "address"},
	}

	userOne, ok := users[userOneId]
	if !ok {
		panic("i tried to check myself and failed, sadge")
	}

	userOne.name = "new name"
	users[userOneId] = userOne
}

type user struct {
	id      int
	name    string
	address string
}

type userId struct {
	id            int
	specialNumber int
}
