package main

import (
	"fmt"
	//"github.com/gin-gonic/gin"
)

func MinMax(y []int, z []int) {

	min := 999999999999999999
	var max = 0

	for _, number := range y {
		if number < min {
			min = number
		}
	}
	fmt.Println("min num: ", min)

	for _, number := range z {
		if number > max {
			max = number
		}
	}
	fmt.Println("max num: ", max)

}

func main() {

	MinMax([]int{8, 12, 56, 4, 3}, []int{8, 12, 56, 4, 3})
}
