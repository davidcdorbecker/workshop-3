package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

const (
	logInfo = "INFO"
)

type log struct {
	severity string
	message  string
}

var (
	logCh  = make(chan log, 50)
	doneCh = make(chan struct{})
)

func main() {
	//ch := make(chan int)
	//
	//wg.Add(1)
	//go func(ch <-chan int) {
	//
	//	for i := range ch {
	//		fmt.Println(i)
	//	}
	//
	//	wg.Done()
	//}(ch)
	//
	//wg.Add(1)
	//go func(ch chan<- int) {
	//	ch <- 42
	//	ch <- 55
	//	ch <- 41
	//	ch <- 50
	//	close(ch)
	//	wg.Done()
	//}(ch)
	//
	//wg.Wait()

	//using select statement and signal channels
	go logger()
	logCh <- log{logInfo, "App is starting"}
	logCh <- log{logInfo, "Shutting down"}
	doneCh <- struct{}{}

	time.Sleep(200 * time.Millisecond)
}

func logger() {

	//for entry := range logCh {
	//	fmt.Printf("%s - %s\n", entry.severity, entry.message)
	//}

	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%s - %s\n", entry.severity, entry.message)
		case <-doneCh:
			fmt.Println("end logger!")
			break
		}
	}
}
