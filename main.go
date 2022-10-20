package main

import (
	"fmt"
	//"github.com/gin-gonic/gin"
)

func Max(x ...int) {

	var max = 0

	for _, number := range x {
		if number > max {
			max = number
		}
	}
	fmt.Println("max num: ", max)
}

func Min(y ...int) {

	min := 999999999999999999

	for _, number := range y {
		if number < min {
			min = number
		}
	}
	fmt.Println("min num: ", min)
}

func main() {

	Min([]int{8, 12, 56, 4, 3}...)
	Max([]int{8, 12, 56, 4, 3}...)

}
