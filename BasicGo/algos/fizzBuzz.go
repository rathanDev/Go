package main

import "fmt"

func main3() {
	findFizzBuzz(100)
}

func findFizzBuzz(n int32) {

	for i := int32(1); i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
