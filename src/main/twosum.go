package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now().UnixNano()
	nums := []int{0, 4, 3, 0}
	index := twoSum2(nums, 0)
	fmt.Println(index)
	end := time.Now().UnixNano()
	fmt.Println((end - start) / 1000)
}

func twoSum(nums []int, target int) []int {

	for w, wVal := range nums {
		for i, iVal := range nums {
			if iVal + wVal == target && i != w {
				return []int{w, i}
			}
		}
	}

	return nil
}


func twoSum2(nums []int, target int, ) []int {

	visited := make(map[int]int)

	for index, val := range nums {
		preIndex, ok := visited[target - val]
		if ok {
			return []int{preIndex, index}
		} else {
			visited[val] = index
		}
	}

	return nil

}
