package main

import "fmt"

/**
切片 切片是对数组的抽象
*/
func main() {

	//null_slice()
	sub_slice()
}

// 空切片
func null_slice() {
	var numbers []int
	printSlice(numbers)
	if numbers == nil {
		fmt.Printf("切片是空的")
	}
}

//切片截取 [:end],[start:],[3:5]
func sub_slice() {
	//创建数组，切片
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	printSlice(numbers)

	//打印原始切片

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

}


//切片追加


func printSlice(numbers []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)
}
