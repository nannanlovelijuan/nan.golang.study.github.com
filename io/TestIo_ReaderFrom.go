package io

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//readFileByReaderFrom()
	function1()
}

/*ReaderFrom 从r中读取数据，知道EOF或发生错误,返回值 n 为读取的字节数*/

func readFileByReaderFrom() {
	file, err := os.Open("D:\\文档管理\\工作情况.md")
	if err != nil {
		fmt.Printf("读取文件出错了")
		panic(err)
	}
	//这个函数调用不是普通的函数调用，而是会在函数正常返回，也就是return之后添加一个函数调用。因此，defer通常用来释放函数内部变量。
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	n, err := writer.ReadFrom(file)

	fmt.Printf("总共读取了%d字节", n)
	writer.Flush()
}

//如果不通过 ReadFrom 接口来做这件事，而是使用 io.Reader 接口，我们有两种思路：
//先获取文件的大小（File 的 Stat 方法），之后定义一个该大小的 []byte，通过 Read 一次性读取
//定义一个小的 []byte，不断的调用 Read 方法直到遇到 EOF，将所有读取到的 []byte 连接到一起

func function1() {
	file, err := os.Open("D:\\文档管理\\工作情况.md")
	if err != nil {
		fmt.Printf("读取文件出错了")
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("获取文件信息出错了")
	}
	length := fileInfo.Size()

	fmt.Printf("file总共的字节数:%d\n", length)

	var p = make([]byte, length)
	n, err := file.Read(p)

	fmt.Printf("file.Read总共读的字节数:%d\n", n)

	fmt.Printf("读取的数据为:%s", string(p))
}

//func function2()  {
//	file, err := os.Open("D:\\文档管理\\工作情况.md")
//	if err != nil {
//		fmt.Printf("读取文件出错了")
//		panic(err)
//	}
//	fileInfo, err := file.Stat()
//	if err != nil {
//		fmt.Printf("获取文件信息出错了")
//	}
//	length := fileInfo.Size()
//
//
//	var a = make([]byte,length)
//	//每次读取100个字节
//	var p = make([]byte, 100)

	//n,err :=file.Read(p)
	//
	//if {
	//
	//}

//}