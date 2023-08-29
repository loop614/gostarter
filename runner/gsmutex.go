package runner

import (
	"fmt"
	"sync"
	"time"
)

func runMutex() {
	printStart()
	defer printEnd()
	mux := sync.Mutex{}
	mux.Lock()
	defer mux.Unlock()
	var wg sync.WaitGroup

	myVariable := 0
	wg.Add(1)
	go changeVariable(&myVariable, &mux, &wg)
	wg.Add(1)
	go changeVariable2(&myVariable, &mux, &wg)
	wg.Wait()
	fmt.Println(myVariable)
}

func changeVariable(i *int, m *sync.Mutex, wg *sync.WaitGroup) {
	defer m.Lock()
	defer wg.Done()
	time.Sleep(time.Second)
	m.Unlock()
	*i++
}

func changeVariable2(i *int, m *sync.Mutex, wg *sync.WaitGroup) {
	defer m.Lock()
	defer wg.Done()
	time.Sleep(time.Second * 2)
	m.Unlock()
	*i++
}
