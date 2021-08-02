package main

import "fmt"

func main() {

	//range_arr_or_slice()
	range_map()
}

//range 切片/数组
func range_arr_or_slice() {
	var numbers = []int{1, 2, 3, 4, 5, 6}

	var sum = 0
	for _, v := range numbers {
		sum += v
	}
	println(sum)
}

// range map

func range_map()  {
	var kvs = map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs{
		fmt.Printf("key=%s,value=%s\n",k,v)
	}
}