package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var counter = 0

var wg = sync.WaitGroup{}

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("after 1 second")
	}()

	////race condition
	//msg := "hello world!"
	//go func() {
	//	fmt.Println(msg)
	//}()
	//msg = "bye!"
	//time.Sleep(time.Second)

	//for i := 0; i < 5; i++ {
	//	wg.Add(2)
	//	go sayHello()
	//	go increment()
	//}
	//
	//wg.Wait()
	//
	//runtime.GOMAXPROCS(100)
}

func increment() {
	counter++
	wg.Done()
}

func sayHello() {
	fmt.Println("hello $", counter)
	wg.Done()
}
