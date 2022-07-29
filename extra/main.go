package main

import (
	"fmt"
	"sync"
	"time"
)

//cond
//func main() {
//	//c := sync.NewCond(&sync.Mutex{})
//	//queue := make([]interface{}, 0, 10)
//	//removeFromQueue := func(delay time.Duration) {
//	//	time.Sleep(delay)
//	//	c.L.Lock()
//	//	queue = queue[1:]
//	//	fmt.Println("Removed from queue")
//	//	c.L.Unlock()
//	//	c.Signal()
//	//}
//	//for i := 0; i < 8; i++ {
//	//	c.L.Lock()
//	//	for len(queue) == 2 {
//	//		c.Wait()
//	//	}
//	//	fmt.Println("Adding to queue")
//	//	queue = append(queue, struct{}{})
//	//	go removeFromQueue(1 * time.Second)
//	//	c.L.Unlock()
//	//}
//
//	type Button struct {
//		Clicked *sync.Cond
//	}
//	button := Button{ Clicked: sync.NewCond(&sync.Mutex{}) }
//	subscribe := func(c *sync.Cond, fn func()) {
//		var goroutineRunning sync.WaitGroup
//		goroutineRunning.Add(1)
//		go func() {
//			goroutineRunning.Done()
//			c.L.Lock()
//			defer c.L.Unlock()
//			c.Wait()
//			fn()
//		}()
//		goroutineRunning.Wait()
//	}
//	var clickRegistered sync.WaitGroup
//
//	clickRegistered.Add(3)
//	subscribe(button.Clicked, func() {
//		fmt.Println("Maximizing window.")
//		clickRegistered.Done()
//	})
//	subscribe(button.Clicked, func() {
//		fmt.Println("Displaying annoying dialog box!")
//		clickRegistered.Done()
//	})
//	subscribe(button.Clicked, func() {
//		fmt.Println("Mouse clicked.")
//		clickRegistered.Done()
//	})
//	button.Clicked.Broadcast()
//	clickRegistered.Wait()
//}

//once
//func main() {
//	var count int
//	increment := func() {
//		count++
//	}
//
//	var once sync.Once
//	var increments sync.WaitGroup
//	increments.Add(100)
//	for i := 0; i < 100; i++ {
//		go func() {
//			defer increments.Done()
//			once.Do(increment)
//		}()
//	}
//	increments.Wait()
//	fmt.Printf("Count is %d\n", count)
//}

//singleton
var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	time.Sleep(time.Second)
}