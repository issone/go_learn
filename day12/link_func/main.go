package main

import "fmt"

type Stu struct {
	Name string
	Age  int
}

func (p *Stu) SetName(name string) *Stu {
	p.Name = name
	return p
}

func (p *Stu) SetAge(age int) *Stu {
	p.Age = age
	return p
}

func (p *Stu) Print() {
	fmt.Printf("age:%d name:%s\n", p.Age, p.Name)
}

// 链式操作，即将结构体指针返回，后续可以通过.操作

func main() {
	stu := &Stu{}
	stu.SetAge(12).SetName("stu01").Print()
	//stu.SetName("stu01")
	//stu.Print()
}
