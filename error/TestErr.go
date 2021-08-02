package main

import "fmt"

func main() {

	// 正常情况
	result, errorMsg := chufa(100, 10)
	if errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	result, er := chufa(100, 0)
	if er != "" {
		fmt.Println("errorMsg is: ", er)
	}
}

//定义除法类

type Division struct {
	fenzi int
	fenmu int
}

func (division *Division) Error() string {

	return "error"
}

func chufa(fenzi int, fenmu int) (result int, error string) {
	if fenmu == 0 {
		var a = Division{fenzi, fenmu}
		error = a.Error()
		return
	} else {
		return fenzi / fenmu, ""
	}
}
