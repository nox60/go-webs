package main

import (
	"fmt"
	"strings"
)

type testStruct struct {
	RoleId    int         `json:"roleId"`
	Code      string      `json:"code"`
	Name      string      `json:"name"`
	Status    int         `json:"status"`
	Functions interface{} `json:"functionIds"`
}

func main() {

	a := ""

	c := strings.Split(a, ",")
	fmt.Println(len(c))

	fmt.Println(c)

	arr := []string{"hello", "world", "1", "2"}

	arrString := strings.Join(arr, ",")

	fmt.Println(arrString)

}
