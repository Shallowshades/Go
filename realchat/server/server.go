package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip        string           //ip地址
	Port      int              //端口
	OnlineMap map[string]*User //在线用户列表
	mapLock   sync.RWMutex     //读写锁
	Message   chan string      //广播消息的channel
}

// 创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线User
func (this *Server) ListenMessager() {

	for {
		msg := <-this.Message

		//将msg发送给全部在线的User
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 当前连接的业务
func (this *Server) Handler(conn net.Conn) {

	user := NewUser(conn, this)
	user.Online()

	//监听用户是否活跃的channel
	isLive := make(chan bool)

	//接受客户端发送的消息,每个User都有一个
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 { //已下线
				user.Offline()
				return
			}
			if err != nil { //读取出错
				fmt.Println("Conn Read error...")
				return
			}

			//提取用户的消息（去除'\n）
			msg := string(buf[:n-1])

			//将得到的消息广播
			user.DoMessage(msg)
			isLive <- true
		}
	}()

	//当前handler阻塞
	for {
		select {
		case <-isLive: //当前用户活跃，重新设置定时器

		case <-time.After(time.Minute * 10): //已经超时，强制关闭连接
			user.SendMsg("long time no say, forced offline")
			close(user.C) //关闭通道
			conn.Close()  //关闭连接
			return        //退出当前handler
		}
	}

}

// 启动服务器接口
func (this *Server) Start() {

	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	//close listen socket
	defer listener.Close()
	//启动监听Message的goroutine
	go this.ListenMessager()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accpte err:", err)
			continue
		}

		//a link coming
		fmt.Println("There is a link...")

		//do handler
		go this.Handler(conn)
	}
}
