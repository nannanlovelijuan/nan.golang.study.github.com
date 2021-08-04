package io

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//func main() {
//
//	//read_stdin()
//
//	//read_file()
//
//	readString()
//}

func readString()  {
	data, err := ReadFrom(strings.NewReader("from string"), 11)

	if err != nil {
		fmt.Println("读取错误")
	}else {
		for i,value := range data{
			fmt.Printf("index:%d,value:%d \n",i,value)
		}
		//将byte[]数组转换为string
		println("将byte[]数组转换为string",string(data))
	}
}
/* 从文件读取*/
func readFile()  {
	file, err := os.Open("D:\\文档管理\\工作情况.md")
	if(err !=nil){
		fmt.Printf("读取文件出错")
	}else {
		data, err:= ReadFrom(file,2048)
		if err != nil {
			fmt.Println("读取错误---->",err)
			return
		}else {
			for i,value := range data{
				fmt.Printf("index:%d,value:%d \n",i,value)
			}
			println("将byte[]数组转换为string",string(data))
		}
	}
	err = file.Close()
	if err != nil {
		fmt.Println("close file error")
	}
}

//从标准输入读取
func read_stdin()  {
	data, err := ReadFrom(os.Stdin, 11)
	if err != nil {
		fmt.Println("读取错误")
	}else {
		for i,value := range data{
			fmt.Printf("index:%d,value:%d \n",i,value)
		}

		//将byte[]数组转换为string
		println("将byte[]数组转换为string",string(data))
	}
}

func ReadFrom(reader io.Reader, num int) ([]byte ,error) {

	var p = make([]byte, num)
	n, err := reader.Read(p)
	if n > 0{
		return p[:n],nil
	}
	return p,err
}


