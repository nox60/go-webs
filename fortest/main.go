package main

import "fmt"

func main() {
	i := 0
	defer fmt.Println("aaa", i) //输出0，因为i此时就是0
	i++
	defer fmt.Println("bbb", i) //输出1，因为i此时就是1
	return
}
