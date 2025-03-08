// https://www.fullstack.cafe/blog/go-interview-questions

package main

import "fmt"

func main() {
	// checkMap()
	// rev2()
	// sliceSlice()

	// total := argSum(3, 5, 1)
	// fmt.Println(total)

	// sum := arraySum([]int32{3,5,7,9})
	// fmt.Print(sum)

	mergedSlice := merge([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)
	fmt.Println("Merged:---->", mergedSlice)
}

func merge(nums1 []int, m int, nums2 []int, n int) []int{
	sl := make([]int, 0)
	i := 0 
	i1 := 0 
	i2 := 0 

	for i1 < m && i2 < n {
		n1 := nums1[i1]
		n2 := nums2[i2]
		if n1 < n2 {
			sl = append(sl, n1)
			i1++
		} else {
			sl = append(sl, n2)
			i2++
		}
		i++
		fmt.Println(sl)
	}

	if i1 < m {
		for i:=i1; i<m; i++ {
			sl = append(sl, nums1[i])
			fmt.Println(sl)
		}
	}

	if i2 < n {
		for i:=i2; i<n; i++ {
			sl = append(sl, nums2[i])
			fmt.Println(sl)
		}
	}

	for i, val := range sl {
		nums1[i] = val
	}

	return nums1
}







func arraySum(ar []int32) int32 {
	var sum int32 = 2
	for _, val := range ar {
		sum += val
	}
	return sum
}

func argSum(nums ...int) int {
	// return nums[0] + nums[1]
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

func sliceSlice() {
	sl1 := []int{1, 2, 3}
	sl2 := []int{4, 5, 6}
	sl1 = append(sl1, sl2...)
	fmt.Println(sl1)
}

func rev() {
	ar := make([]int, 2)
	// ar := []int{1, 2, 3}
	ar[0] = 1
	ar[1] = 2
	ar[2] = 3
	fmt.Println(ar)
	// error: Index out of range
}

func rev2() {
	// var slice []int					// start with an empty slice
	slice := make([]int, 0, 2) // start slice with capacity 2

	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)
	fmt.Println(slice)
}

func checkMap() {
	oneMap := make(map[string]string)
	fmt.Println(oneMap)

	oneMap["key1"] = "val1"

	val, ok := oneMap["key1"]
	fmt.Println(val, ok)
	if ok {
		fmt.Println("Val1 exists")
	}

	if val, ok := oneMap["key1"]; ok {
		fmt.Println("Here", val)
	}
}
