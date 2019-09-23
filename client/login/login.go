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
		UserID:  100,
		UserPwd: "200",
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
	return nil
}
