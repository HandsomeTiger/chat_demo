package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

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
	m, err := readPkg(conn)
	if err != nil {
		fmt.Println("read pkg failed ", err.Error())
		return
	}
	err = serverProcessMsg(conn, m)
	if err != nil {
		fmt.Println("server process failed ", err.Error())
		return
	}
	fmt.Println("process success")
}

func readPkg(conn net.Conn) (*message.Message, error) {
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
		return &req, nil
	}
}

func serverProcessMsg(conn net.Conn, msg *message.Message) error {
	switch msg.Type {
	case message.LoginMessgeType:
		// 处理登录
		serverProcessLogin(conn, msg)
	case message.RegisterMessgeType:
	// 处理注册

	default:
		return errors.New("类型不合法")
	}
	return nil
}

// serverProcessLogin
func serverProcessLogin(conn net.Conn, msg *message.Message) error {
	loginMsg := &message.LoginMessage{}
	err := json.Unmarshal([]byte(msg.Data), loginMsg)
	if err != nil {
		return err
	}
	var resMsg message.LoginResponse
	if loginMsg.UserID == 1 && loginMsg.UserPwd == "1" {
		resMsg = message.LoginResponse{
			Code:  200,
			Error: "登录成功",
		}

	} else {
		resMsg = message.LoginResponse{
			Code:  401,
			Error: "用户不合法",
		}
	}
	jres, err := json.Marshal(resMsg)
	if err != nil {
		return err
	}
	fmt.Println(string(jres))
	time.Sleep(2 * time.Second)
	_, err = conn.Write(jres)
	if err != nil {
		return err
	}
	return nil
}
