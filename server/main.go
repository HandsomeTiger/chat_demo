package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器监听8889端口！")
	ls, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("listen failed :", err.Error())
		return
	}
	defer ls.Close()
	for {
		fmt.Println("等待客户端连接")
		conn, err := ls.Accept()
		if err != nil {
			fmt.Println("accept failed:", err.Error())
		}
		go process(conn)
	}
	fmt.Println("关闭服务器")
}

func process(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 8096)
	for {
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("读到buf =", buf[:4])
		fmt.Println("读到buflen =", buf[:4])
	}
}
