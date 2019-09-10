package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"reflect"
)

func main(){
	dial()
}

func dial(){
	conn,err:=redis.Dial("tcp","127.0.0.1:6379")
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer conn.Close()
	fmt.Println(conn)
	_,err=conn.Do("set","a","abs")
	if err!=nil{
		fmt.Printf("操作错误:%v",err.Error())
		return
	}
	fmt.Println("操作成功")
	a,err:=conn.Do("get","a")
	if err!=nil{
		fmt.Printf("读取错误:%v",err.Error())
		return
	}
	fmt.Println(reflect.TypeOf(a))
	r,err:=redis.String(a,err)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("读取成功")
	fmt.Printf("%s",r)
}
