package main

import "fmt"

type Person struct {
	id int
	name string
}

func main() {
	var person1 = Person{1,"abc"}
	println(person1.name)
	fmt.Printf("&person1,类型为%T\n",&person1)
	fmt.Printf("&person1,内存地址为%p\n",&person1)
	fmt.Printf("person1,类型为%T\n",person1)
	println_person(&person1)
}

func println_person(person *Person)  {
	fmt.Printf("方法内部类型为%T\n",person)
	fmt.Printf("方法内部内存地址为%p\n",person)
	fmt.Println(person)
	fmt.Println(person.id)
}