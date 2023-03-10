package main

import (
	"ZinxV1.0/ziface"
	"ZinxV1.0/znet"
	"fmt"
)

// PingRouter ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Handle Ping
func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("received from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendBuffMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

// Handle HelloZinxRouter
func (h *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("received from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendBuffMsg(1, []byte("Hello Zinx Router V1.0"))
	if err != nil {
		fmt.Println(err)
	}
}

// DoConnectionBegin 创建连接的时候执行
func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("DoConnectionBegin is Called ... ")

	//设置两个链接属性，在连接创建之后
	fmt.Println("Set conn Name, Home done!")
	conn.SetProperty("Name", "HerrKung")
	conn.SetProperty("Home", "HerrKung@outlook.com")

	err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	if err != nil {
		fmt.Println(err)
	}
}

// DoConnectionLost 连接断开的时候执行
func DoConnectionLost(conn ziface.IConnection) {
	//在连接销毁之前，查询conn的Name，Home属性
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Conn Property Name = ", name)
	}

	if home, err := conn.GetProperty("Home"); err == nil {
		fmt.Println("Conn Property Home = ", home)
	}

	fmt.Println("DoConnectionLost is Called ... ")
}

func main() {
	//创建一个server句柄
	s := znet.NewServer()

	//注册链接hook回调函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	//配置路由
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	//开启服务
	s.Serve()
}
