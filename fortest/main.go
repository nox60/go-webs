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
	//i := 0
	//defer fmt.Println("aaa", i) //输出0，因为i此时就是0
	//i++
	//defer fmt.Println("bbb", i) //输出1，因为i此时就是1

	var s1 []int64

	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 33)
	s1 = append(s1, 21)
	s1 = append(s1, 24)

	for index, value := range s1 {
		fmt.Println("index:", index, "value:", value)
	}

	return
}
