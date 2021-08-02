package main

import "fmt"

func main() {

}

//标准for循环
func for1() {
	var sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
}

//简写
func for2() {
	var sum = 0
	for ; sum < 10; {
		sum += sum
	}
}

//range

func for3() {
	var strings = []string{"a", "b"}
	for i, s:=range strings{
		fmt.Println(i,s)
	}

	var ints = []int{1,2,3,4}
	for i, s:=range ints{
		fmt.Printf("第%d个为%d",i,s)
	}
}
