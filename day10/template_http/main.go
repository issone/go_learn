package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var myTemplate *template.Template

type Result struct {
	output string
}

func (p *Result) Write(b []byte) (n int, err error) {
	fmt.Println("called by template")
	p.output += string(b)
	return len(b), nil
}

type Person struct {
	Name  string
	Title string
	Age   int
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	//fmt.Fprintf(w, "hello ")
	var arr []Person
	p := Person{Name: "Mary001", Age: 10, Title: "我的个人网站"}
	p1 := Person{Name: "Mary002", Age: 11, Title: "我的个人网站"}
	p2 := Person{Name: "Mary003", Age: 12, Title: "我的个人网站"}
	arr = append(arr, p)
	arr = append(arr, p1)
	arr = append(arr, p2)

	resultWriter := &Result{}
	io.WriteString(resultWriter, "hello world")
	err := myTemplate.Execute(w, arr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("template render data:", resultWriter.output)

}

func initTemplate(filename string) (err error) {
	myTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}

func main() {
	initTemplate("E:/go_projects/src/go_learn/day10/template_http/index.html")
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
