package two

import (
	"fmt"
	"time"
)

func xmain() {
	fmt.Println("Starts")

	c := make(chan string)

	go count("fish", c)

	for {
		msg, open := <-c // BlocksOrWaits to receive msg from channel
		if !open {
			break
		}
		fmt.Println(msg)
	}

	fmt.Println("Ends")
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Second * 1)
	}
	close(c)
}

// func main() {
// 	fmt.Println("Starts main")

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go count("fish", 1, &wg)
// 	go count("sheep", 2, &wg)

// 	wg.Wait()

// 	fmt.Println("End")
// }

// func count(thing string, duration time.Duration, wg *sync.WaitGroup) {
// 	for i:=1; i<=5; i++ {
// 		fmt.Println(i, thing)
// 		time.Sleep(time.Second * duration)
// 	}
// 	wg.Done()
// }



// https://www.fullstack.cafe/blog/go-interview-questions