package main

import (
	"fmt"

	"github.com/handsomeTiger/chat_demo/client/login"
)

// 定义两个变量，一个表示用户id，一个表示用户密码
var userID int
var userPwd string

func main() {
	var key int
	var loop bool = true
	for loop {
		fmt.Println("-----------欢迎登陆多人聊天系统------------")
		fmt.Println("\t\t\t 1. 登录聊天室")
		fmt.Println("\t\t\t 2. 注册用户")
		fmt.Println("\t\t\t 3. 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("开始登录")
			loop = false
		case 2:
			fmt.Println("开始注册")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}

	if key == 1 {
		fmt.Println("请输入用户的id")
		fmt.Scanf("%d\n", &userID)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%s\n", &userPwd)
		// login
		err := login.Login(userID, userPwd)
		if err != nil {
			fmt.Println(err.Error())
			return
		} else {
			fmt.Println("登录成功")
		}
	} else if key == 2 {

	} else {

	}
}
