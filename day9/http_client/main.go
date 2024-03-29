package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn,err := net.Dial("tcp","www.baidu.com:80")
	if err !=nil{
		fmt.Println("Error dialing",err.Error())
		return
	}
	defer conn.Close()

	msg := "GET / HTTP/1.1\r\n"
	msg += "Host: www.baidu.com\r\n"
	msg += "Connection: keep-alive\r\n"
	msg += "\r\n\r\n"

	//io.WriteString(os.Stdout, msg)
	n, err := io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("write string failed, ", err)
		return
	}
	fmt.Println("send to baidu.com bytes:", n)
	buf := make([]byte, 4096)
	for {
		count, err := conn.Read(buf)
		fmt.Println("count:", count, "err:", err)
		if err != nil {
			break
		}
		fmt.Println(string(buf[:count]))
	}
}
