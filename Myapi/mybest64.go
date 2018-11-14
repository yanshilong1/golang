package Myapi

import (
	b64 "encoding/base64"
	"fmt"
)

func Mybaes64() {

	data := "abc!@#$dsgrsbv_*@  343"
	sEn := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("sen: ", sEn)

	sDc, _ := b64.StdEncoding.DecodeString(sEn)
	fmt.Println("sdc: ", string(sDc))

	uEn := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("yen: ", uEn)
	uDn, _ := b64.URLEncoding.DecodeString(string(uEn))
	fmt.Println("udn: ", uDn)

}
