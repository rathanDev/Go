package main

import (
	"fmt"
	"sync"
)

func mainC() {
	ch := make(chan string)
	go getHello(ch)
	go sayHello(ch)
	select {}
}

func sayHello(ch chan string) {
	ch <- "Hi hi!"
	close(ch)
}

func getHello(ch chan string) {
	for val := range ch {
		fmt.Println(val)
	}
}









func main2() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go sayHello2(&wg)

	// time.Sleep(time.Second * 3)

	wg.Wait()

}

func sayHello2(wg *sync.WaitGroup) {
	// fmt.Println(&wg)
	fmt.Println("Hey!")
	wg.Done()
}





func main1() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello1(&wg)
	wg.Wait()
}

func sayHello1(wg *sync.WaitGroup) {
	fmt.Println("Hey!")
	wg.Done()
}