package Myinterface

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokaiaphone NokiaPhone) call() {

	fmt.Println("my name is nokia iam  calling you ")

}

type Iphone struct {
}

func (iphone Iphone) call() {
	fmt.Println("i am a phone im calling you ")

}

type TTTT struct {
}

func (ttt TTTT) call() {
	fmt.Println("this is my calling -----")
	fmt.Println("请输入您的参数:")
	var a int
	fmt.Scan(&a)
	fmt.Println(a)

}

func Dointerface() {

	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(Iphone)
	phone.call()

	phone = new(TTTT)
	phone.call()

}
