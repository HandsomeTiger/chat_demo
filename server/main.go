package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/handsomeTiger/chat_demo/common/message"
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
	readPkg(conn)
}

func readPkg(conn net.Conn) (*message.LoginMessage, error) {
	buf := make([]byte, 8096)
	for {
		_, err := conn.Read(buf[:4])
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		var pkhLen uint32 = binary.BigEndian.Uint32(buf[0:4])
		fmt.Printf("read len %d \n", pkhLen)

		n, err := conn.Read(buf[:pkhLen])
		fmt.Println("n=", n)
		if n != int(pkhLen) || err != nil {
			fmt.Printf("err :%v \n", err.Error())
			return nil, err
		}
		fmt.Println(string(buf[:pkhLen]))
		var req message.Message
		if err := json.Unmarshal(buf[:pkhLen], &req); err != nil {
			fmt.Printf("json unmarshal failed %v", err.Error())
			return nil, err
		}
		fmt.Println(req.Data)
		var data message.LoginMessage
		if err := json.Unmarshal([]byte(req.Data), &data); err != nil {
			fmt.Println("err = ", err.Error())
		}

		return &data, nil

	}
}
