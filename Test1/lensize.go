package Test1

import "unsafe"

func GetSize() {
	const (
		a = "abcc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)

	println(a, b, c)

}
