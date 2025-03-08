package two

import "fmt"

func main2() {
	fmt.Println("Run")

	// var x int 
	// x = 10; 
	// fmt.Println(*&x)

	// y := 15
	// fmt.Println(*&y)

	// const z = 3.14
	// fmt.Println(z)

	ar := [3]int{1, 2}
	fmt.Println(ar)

	sl := make([]int, 5)
	fmt.Println(sl)

	sl[0] = 11
	fmt.Println(sl)

	sl = append(sl, 10)
	fmt.Println(sl)



	numMap := make(map[int]int, 5)
	fmt.Println(numMap)
	numMap[5] = 6
	numMap[1] = 2
	fmt.Println(numMap)

}