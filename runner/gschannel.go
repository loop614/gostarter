package runner

import (
	"fmt"
	"time"
)

func runChannel() {
	printStart()
	defer printEnd()

	ch := make(chan chanResult)
	go waitSomeAction(ch)
	go waitSomeAction2(ch)
	result1 := <-ch
	if result1.Error != nil {
		panic("imagine other scenario where this makes sense")
	}
	fmt.Printf("read from channel = %d\n", result1.Result)
	result2 := <-ch
	if result2.Error != nil {
		panic("imagine other scenario where this makes sense")
	}
	fmt.Printf("read from channel = %d\n", result2.Result)
}

type chanResult struct {
	Result int
	Error  error
}

func waitSomeAction(ch chan chanResult) {
	time.Sleep(time.Second * 1)
	ch <- chanResult{1, nil}
}

func waitSomeAction2(ch chan chanResult) {
	time.Sleep(time.Second * 2)
	ch <- chanResult{2, nil}
}
