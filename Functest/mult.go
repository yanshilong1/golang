package Functest

import (
	"fmt"
	"strconv"
)

func Mult() {

	for i := 1; i < 10; i++ {

		for j := 1; j <= i; j++ {
			var ret string
			ret = strconv.Itoa(i * j)
			//字符串转int：strconv.Atoi()
			//int转字符串: strconv.Itoa()
			fmt.Print("  ", i, " * ", j, "=", ret)
		}

		fmt.Println()

	}

}
