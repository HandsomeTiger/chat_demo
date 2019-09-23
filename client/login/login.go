package login

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/handsomeTiger/chat_demo/common/message"
)

func Login(uid int, pwd string) error {
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("dial failed:", err.Error())
		return err
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.LoginMessgeType
	data := message.LoginMessage{
		UserID:  uid,
		UserPwd: pwd,
	}
	encodeDataByte, err := json.Marshal(data)
	if err != nil {
		fmt.Println("marshal failed:", err.Error())
		return err
	}
	mes.Data = string(encodeDataByte)
	encodeMes, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("mes marshal failed", err.Error())
		return err
	}
	// int转 []byte
	var pkgLen uint32 = uint32(len(encodeMes))
	fmt.Println("len = ", pkgLen)
	var lenByte [4]byte
	binary.BigEndian.PutUint32(lenByte[0:4], pkgLen)
	n, err := conn.Write(lenByte[0:4])
	if n != 4 || err != nil {
		fmt.Println("write failed")
		return err
	}
	fmt.Println("客户端发送长度成功 len ")
	time.Sleep(1 * time.Second)
	fmt.Println(string(encodeMes))
	fmt.Println(mes)
	_, err = conn.Write(encodeMes)
	if err != nil {
		fmt.Println("write data failed err = ", err.Error())
		return err
	}
	fmt.Println("等待响应")
	res := &message.LoginResponse{}

	var resBytes = make([]byte, 9000)
	n, err = conn.Read(resBytes)
	if err != nil {
		fmt.Println("read 失败", err.Error())
		return err
	}
	fmt.Println(string(resBytes))
	err = json.Unmarshal(resBytes[:n], res)
	if err != nil {
		return err
	}
	fmt.Println(res.Code)
	if res.Code == 200 {
		fmt.Println("登录成功")
		fmt.Println(res.Error)
	} else {
		fmt.Println("登录失败")
		fmt.Println(res.Error)
	}
	return nil
}
