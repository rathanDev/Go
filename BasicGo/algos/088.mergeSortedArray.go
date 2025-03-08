package main

import "fmt"

// https://leetcode.com/problems/merge-sorted-array/submissions/1566863380/?envType=study-plan-v2&envId=top-interview-150

func main() {
	// merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	// merge([]int{0}, 0, []int{1}, 1)
	merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	i := m + n - 1
	i1 := m - 1
	i2 := n - 1

	for i1 >= 0 && i2 >= 0 {
		n1 := nums1[i1]
		n2 := nums2[i2]
		if n1 > n2 {
			nums1[i] = n1
			i1--
		} else {
			nums1[i] = n2
			i2--
		}
		i--
	}

	if i2 >= 0 {
		for i2 >= 0 {
			nums1[i] = nums2[i2]
			i2--
			i--
		}
	}

	fmt.Println(nums1)

}
