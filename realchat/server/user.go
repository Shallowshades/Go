package main

import (
	"fmt"
	"net"
	"runtime"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	//监听user channel的goroutine
	go user.ListenMessage()

	return user
}

// 用户上线
func (this *User) Online() {
	//将用户加入到onlinemap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	//广播用户上线
	this.server.BroadCast(this, "online...")
}

// 用户下线
func (this *User) Offline() {
	//将用户从onlinemap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
	//广播用户下线
	this.server.BroadCast(this, "outline...")
}

// 给当前User客户端发送消息
func (this *User) SendMsg(msg string) {
	_, err := this.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("write error:", err)
		runtime.Goexit()
	}
}

// 用户处理消息业务
func (this *User) DoMessage(msg string) {
	if msg[0:3] == "who" {
		//查询在线用户都有哪些
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			if user != this {
				onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "is online ...\n"
				this.SendMsg(onlineMsg)
			}
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename " {
		//重命名 格式：rename 张三
		newName := strings.Split(msg, " ")[1]
		//判断name是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok { //用户名已存在
			this.SendMsg("This name had occupied\n")
		} else { //用户名不存在
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("Rename is successful!\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to " {
		//私聊 格式to name 消息内容
		//1 获取用户名
		remoteName := strings.Split(msg, " ")[1]
		if remoteName == "" {
			this.SendMsg("format error, format:to name msg\n")
			return
		}
		//2 根据用户名 得到对方User对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("User isn't exist\n")
			return
		}
		//3 获取消息内容，通过对方的User对象将消息内容发送过去
		content := strings.Split(msg, " ")[2]
		if content == "" {
			this.SendMsg("no msg, please resend")
			return
		}
		remoteUser.SendMsg(this.Name + " say: " + content + "\n")

	} else {
		this.server.BroadCast(this, msg)
	}
}

// 监听当前User channel的方法，一旦有消息就直接发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

// // 查询用户是否存在
// func (this *User) Exist(userName string) string {
// 	this.server.mapLock.Lock()
// 	defer this.server.mapLock.Unlock()
// 	for _, user := range this.server.OnlineMap {
// 		if user.Name == userName {
// 			return "true"
// 		}
// 	}
// 	return "false"
// }
