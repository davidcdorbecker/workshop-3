package main

import (
	"fmt"
	"runtime"
	"sync"
)

var m = sync.RWMutex{}
var counter = 0

var wg = sync.WaitGroup{}

func main() {
	//IIFE
	//wg.Add(2)
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	fmt.Println("GR1 after 1 second")
	//	wg.Done()
	//}()
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	fmt.Println("GR2 after 1 second")
	//	wg.Done()
	//}()
	//wg.Wait()

	////race condition
	//msg := "hello world!"
	//go func(msg string) {
	//	fmt.Println(msg)
	//}(msg)
	//msg = "bye!"
	//time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		wg.Add(2)
		m.Lock()
		go sayHello()
		m.Unlock()
		m.RLock()
		go increment()
		m.RUnlock()
	}
	//
	wg.Wait()
	//
	runtime.GOMAXPROCS(-1)
}

func increment() {
	counter++

	wg.Done()
}

func sayHello() {
	fmt.Println("hello $", counter)

	wg.Done()
}
