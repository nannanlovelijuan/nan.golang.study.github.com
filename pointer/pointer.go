package main

import (
	"fmt"
)

//变量、指针和地址三者的关系是，每个变量都拥有地址，指针的值就是地址

func main() {

	//printPointerAdress("a", "b")
	//getValueFromPointer("abc")
	//var x,y = 1,2
	//swap(&x,&y)
	//fmt.Println(x,y)

	mySalary := 80000
	fmt.Printf("变量 mySalary 的内存地址为：%p\n", &mySalary)
	modifySalary(mySalary)
	fmt.Println(mySalary)

	// 指针类型参数，实参的值被改变
	modifySalary2(&mySalary)
	fmt.Println(mySalary)
}

/**
认识指针地址和指针类型
使用&符号获取变量的内存地址,%p 指针占位符
var ptr = &v   v 代表被取地址的变量，变量 v 的地址使用变量 ptr 进行接收，ptr 的类型为*T，称做 T 的指针类型，*代表指针
*/
func printPointerAdress(a, b string) {
	fmt.Printf("%p, %p", &a, &b)
}

//从指针获取指针指向的值
func getValueFromPointer(a string) {
	var temp = &a
	//打印类型
	fmt.Printf("类型：%T\n", temp)
	//打印内存地址
	fmt.Printf("内存地址:%p\n", temp)
	//从指针取值
	var value = *temp
	//打印值的类型
	fmt.Printf("value 类型:%T\n", value)
	//打印值
	fmt.Printf("value: %s", value)
}

func modifySalary (salary int) {
	fmt.Printf("参数变量的内存地址为：%p\n", &salary)
	salary = 100000
}

func modifySalary2 (salary *int) {
	fmt.Printf("参数变量的内存地址为：%p\n", salary)
	*salary = 100000
}
//使用指针修改值
func swap(a,b *int)  {
	//从指针取值
	var t = *a
	//取b指针的值，赋给a指针指向的变量
	*a = *b
	*b = t
}
