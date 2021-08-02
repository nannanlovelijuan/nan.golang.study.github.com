package main

import "fmt"

func main() {
	//make_map()
	map_exist_key()
}

//创建map

func make_map() {
	var map1 map[string]string
	map1 = make(map[string]string)

	map1["a"] = "b"
	map1["c"] = "d"

	for a := range map1 {
		fmt.Println(a, "的值是：", map1[a])
	}
}

// 查看元素是否在map中存在
func map_exist_key() {
	var map1 map[string]string

	map1 = make(map[string]string)

	map1["a"] = "b"
	map1["b"] = "c"

	value,ok := map1["b"]

	if(ok){
		fmt.Printf("存在key为b的值%s",value)
	}else {
		fmt.Println("不存在key为b的值")
	}
}
