package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	n := 0
	for y := 0; y < 10000; y++ {
		go func() {
			m.Lock()
			n++
			m.Unlock()
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println(n)
}

func ChangeNumber(n *int, value int) {
	*n = value
}
