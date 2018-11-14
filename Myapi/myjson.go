package Myapi

import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Reponse2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func Myjson() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	strB, _ := json.Marshal("wewewgfgbgeryydf")
	fmt.Println(string(strB))

	slD := []string{"tesy2", "yset2", "hssr3"}
	slB, _ := json.Marshal(slD)
	fmt.Println("字符串数组转json", slB)

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println("map转json", string(mapB))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println("结构体转数组: ", string(res1B))

}
