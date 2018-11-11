package Functest

import "fmt"

const LENGTH = 10

func Mytaiangle() {

	nums := []int{}
	for i := 0; i < LENGTH; i++ {

		for j := 0; j < (LENGTH - i); j++ {
			fmt.Print(" ")
		}

		for j := 0; j < (i + 1); j++ {
			var length = len(nums)
			var value int

			if j == 0 || j == i {
				value = 1
			} else {
				value = nums[length-i] + nums[length-i-1]
			}

			nums = append(nums, value)
			fmt.Print(value, " ")

		}
		fmt.Println("")
		//

	}

}
