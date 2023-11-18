package main

import (
	"fmt"
)

func main () {
	nums := []int{}
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
		if(i % 2 == 0) {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
	}
}