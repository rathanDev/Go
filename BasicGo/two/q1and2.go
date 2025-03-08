package two

import (
	"fmt"
	"sync"
	"time"
)

// https://medium.com/@ninucium/go-interview-questions-part-1-pointers-channels-and-range-67c61345cf3c

func q1() {
  ch := make(chan *int, 4)
  array := []int{1, 2, 3, 4}
  wg := sync.WaitGroup{}
  wg.Add(1)
  go func() {
    for _, value := range array {
      ch <- &value
    }
	close(ch)
  }()
  go func() {
	
    for value := range ch {
	  time.Sleep(10 * time.Millisecond)
      fmt.Println(*value) // what will be printed here?
    }
    wg.Done()
  }()

  wg.Wait()
}

func q2() {
	array := []int{1, 2, 3, 4}
	newArray := make([]*int, 4)
	for i, value := range array {
	  newArray[i] = &value
	}
	for _, value := range newArray {
	  fmt.Println(*value)
	}
}