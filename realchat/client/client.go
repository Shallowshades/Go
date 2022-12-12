package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端连接
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}
	//链接服务器
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.conn = conn
	return client
}

// 处理server回应的消息，直接显示到标准输出
func (client *Client) DealResponse() {
	//一旦有数据，就直接copy到stdout标准输出上，永久阻塞监听
	for {
		io.Copy(os.Stdout, client.conn)
	}
}

func (client *Client) menu() bool {
	var flag int
	fmt.Println("1. public chat")
	fmt.Println("2. private chat")
	fmt.Println("3. update name")
	fmt.Println("0. exit")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("please input suitable number")
		return false
	}
}

// 更新用户名
func (client *Client) UpdateName() bool {
	fmt.Println("please input new name : ")
	fmt.Scanln(&client.Name)
	if client.Name == "exit" {
		defer client.UpdateName()
		defer fmt.Println("please input suitable name")
		return false

	}
	sendMsg := "rename " + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
}

// 公聊模式
func (client *Client) PublicChat() {
	//提示用户输入信息
	var chatMsg string
	fmt.Println("please input chat content, input exit to exit")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		//发送给服务器

		//消息不为空则发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn.Write err:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println("please input chat content, input exit to exit")
		fmt.Scanln(&chatMsg)
	}
}

// 查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
}

// 查询聊天对象是否在线
func (client *Client) Exist(userName string) bool {
	sendMsg := "exist " + userName
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
}

// 私聊模式
func (client *Client) PrivateChat() {
	var remoteName string
	for {
		client.SelectUsers()
		fmt.Println("input the chat object")
		fmt.Scanln(&remoteName)

		//此时输入任何字符都会进入私聊，但user不一定存在

		if remoteName != "exit" {
			var chatMsg string
			fmt.Println("please input chat content, key exit to exit")
			fmt.Scanln(&chatMsg)

			for chatMsg != "exit" {
				//消息不为空则发送
				if len(chatMsg) != 0 {
					sendMsg := "to " + remoteName + " " + chatMsg + "\n\n"
					_, err := client.conn.Write([]byte(sendMsg))
					if err != nil {
						fmt.Println("conn.Write error:", err)
						return
					}
				}

				chatMsg = ""
				fmt.Scanln(&chatMsg)
			}
		} else {
			return
		}
	}
}

// 业务
func (client *Client) Run() {

	//输入为0时退出
	for client.flag != 0 {
		//等待正确的输入
		for client.menu() != true {
		}

		//根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			client.PublicChat()
		case 2:
			client.PrivateChat()
		case 3:
			client.UpdateName()
		}
	}
}

var serverIp string
var severPort int

//./client -ip 127.0.0.1 -port 8888

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "set ip, default address is 127.0.0.1")
	flag.IntVar(&severPort, "port", 8888, "set port, default port is 8888")
}

func main() {

	//命令行解析
	flag.Parse()

	client := NewClient(serverIp, severPort)
	if client == nil {
		fmt.Println("fail to link server ")
		return
	}
	fmt.Println("succeed to link server")

	//开启一个接受消息的协程
	go client.DealResponse()

	//客户端业务
	client.Run()
}
