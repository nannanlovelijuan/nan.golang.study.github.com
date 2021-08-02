package main

import "fmt"

func main() {

	var phone = new(XiaoMiPhone)
	var phone1 = new(HuaWei)

	phone.call()
	phone1.call()
}

type Phone interface {
	call()
}

type XiaoMiPhone struct {

}

// 实现call接口
func (phone XiaoMiPhone) call() {
	fmt.Println("This is xiao mi call")
}

type HuaWei struct {

}


//实现call接口
func (phone HuaWei) call()  {
	fmt.Println("This is Huawei call")
}