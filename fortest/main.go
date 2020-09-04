package main

import "fmt"

type testStruct struct {
	RoleId    int         `json:"roleId"`
	Code      string      `json:"code"`
	Name      string      `json:"name"`
	Status    int         `json:"status"`
	Functions interface{} `json:"functionIds"`
}

func main() {
	i := 0
	defer fmt.Println("aaa", i) //输出0，因为i此时就是0
	i++
	defer fmt.Println("bbb", i) //输出1，因为i此时就是1
	return
}
