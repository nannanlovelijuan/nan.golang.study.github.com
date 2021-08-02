package main

import "fmt"

func main() {
	const a = 1

	const (
		un_known  = 0
		female = 1
		male = 2
	)
	fmt.Printf("const:%d" , un_known)
	fmt.Printf("const:%d" , female)
	fmt.Printf("const:%d" , male)
}
