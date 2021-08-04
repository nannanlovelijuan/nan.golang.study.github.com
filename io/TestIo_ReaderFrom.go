package io

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//readFileByReaderFrom()
	Function2()
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

func Function2() int {
	file, err1 := os.Open("D:\\文档管理\\工作情况.md")
	f, err2 := os.Create("D:\\文档管理\\工作情况备份.md")
	if err1 != nil {
		fmt.Printf("读取文件出错了")
		panic(err1)
	}
	//每次读取100个字节
	defer file.Close()
	if err2 != nil {
		fmt.Printf("创建文件出错了")
		panic(err2)
	}
	defer f.Close()
	r := bufio.NewReader(file)

	buffer := make([]byte, 1024)
	bytesCount := 0
	for {
		n, err := r.Read(buffer)
		buffer = append(buffer[:n])
		bytesCount += n

		// 控制条件,根据实际调整
		if err != nil && err != io.EOF {
			log.Println(err)
		}
		if n == 0 {
			break
		}

		f.Write(buffer[:n])

	}
	fmt.Printf("------>file.Read总共读的字节数:%d\n", bytesCount)
	return bytesCount
}

func ReadLine() int {
	file, err := os.Open("D:\\文档管理\\工作情况.md")
	if err != nil {
		fmt.Printf("Open file err...%s", err)
		os.Exit(1)
	}
	defer file.Close()

	f,_ := os.Create("D:\\文档管理\\工作情况2.md")

	defer file.Close()
	buff := bufio.NewReader(file)
	bytesCount := 0
	for {
		line, err := buff.ReadBytes('\n')

		//ReadBytes 从输入中读取直到遇到界定符（delim）为止，返回的 slice 包含了从当前到界定符的内容 （包括界定符）
		fmt.Printf("the line:%s,字节大小%d", line,len(line))
		if err != nil && err != io.EOF{
			fmt.Printf("err:%s\n", err)
		}
		if len(line) ==0{
			break
		}
		f.Write(line)
		bytesCount += len(line)
	}

	// 这里可以换上任意的 bufio 的 Read/Write 操作
	return bytesCount
}

func Cat(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}
