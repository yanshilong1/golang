package Struct

import "fmt"

type Books struct {
	titile string
	author string
	age    string
}

func Mystruct1() {

	fmt.Println(Books{"aaaa", "asasa", "2332"})
	fmt.Println(Books{titile: "的标题"})

}
