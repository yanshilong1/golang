package Array

import "fmt"

func Myarray1() {
	var array1 [10]float32
	var array2 = [3]int{1, 2, 3}
	var array3 = [...]float32{11, 2, 3, 1, 4, 2, 4, 2, 24, 421}
	array1[0] = 10
	var a1 int = array2[1]
	var a3 = array3[2]

	fmt.Println(a1, a3)

	for j := 0; j < len(array3); j++ {

		fmt.Println(array3[j])

	}

}
