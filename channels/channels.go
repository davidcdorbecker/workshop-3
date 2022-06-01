package main

import (
	"fmt"
	"sync"
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
	ch := make(chan int)

	wg.Add(2)
	go func() {
		i := <- ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42
		wg.Done()
	}()

	wg.Wait()

	//
	//go func(ch chan <- int) {
	//	ch <- 5
	//	ch <- 10
	//	ch <- 12
	//	close(ch)
	//	wg.Done()
	//}(ch)
	//
	//wg.Wait()

	//using select statement and signal channels
	//go logger()
	//defer func() {
	//	close(logCh)
	//}()
	//logCh <- log{logInfo, "App is starting"}
	//logCh <- log{logInfo, "Shutting down"}
	//doneCh <- struct{}{}
	//
	//time.Sleep(200 * time.Millisecond)
	//logCh <- log{logInfo, "Shutting down"}
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
