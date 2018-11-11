package Struct

import "fmt"

type Books1 struct {
	title string
	name  string
	ang   string
}

func Mstruct2() {

	var book1 Books1
	var book2 Books1

	book1.name = "名字"
	book1.ang = "12"
	book1.title = "标题"

	book2.title = "标题"
	book2.ang = "1212"

	fmt.Println(book2.ang)

	fmt.Println(book2)
	fmt.Println(book1)

}
