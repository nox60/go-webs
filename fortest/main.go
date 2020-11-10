package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
)

type testStruct struct {
	RoleId    int         `json:"roleId"`
	Code      string      `json:"code"`
	Name      string      `json:"name"`
	Status    int         `json:"status"`
	Functions interface{} `json:"functionIds"`
}

func main() {
	//
	//a := ""
	//
	//c := strings.Split(a, ",")
	//fmt.Println(len(c))
	//
	//fmt.Println(c)
	//
	//arr := []string{"hello", "world", "1", "2"}
	//
	//arrString := strings.Join(arr, ",")
	//
	//fmt.Println(arrString)

	for i := 0; i < 10; i++ {
		randonStr := CreateRandomString(8)
		fmt.Println(randonStr)
	}

}

//func main() {
//	randomStr := CreateRandomString(15)
//	ffmt.P(randomStr)
//	//string("mCvYEd8MH8xnBRn")
//}

func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
