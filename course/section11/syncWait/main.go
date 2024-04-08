package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Wait Group
	var wg sync.WaitGroup
	wg.Add(3)
	go ConnectWithDb(&wg)
	go CallApi(&wg)
	go InternalProcess(&wg)

	wg.Wait()
}

func ConnectWithDb(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Db")
	wg.Done()
}

func CallApi(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("API")
	wg.Done()
}

func InternalProcess(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Process")
	wg.Done()
}
