package main

import (
	"fmt"
	"net"
	"time"
)

func Read(socket *net.UDPConn) {

	for {
		recvMsg := make([]byte, 512)
		n, addr, err := socket.ReadFromUDP(recvMsg)
		if err != nil {
			fmt.Println("fail to read, err : ", err)
			return
		}
		fmt.Printf("addr[%v], count(%v), msg : %s\n", addr, n, string(recvMsg[:n]))
	}
}

func main() {

	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("fail to dial udp, err : ", err)
		return
	}
	defer socket.Close()

	//开启读go程
	go Read(socket)

	sendMsg := []byte("Hello, Hello.")
	for i := 0; i < 20; i++ {
		_, err = socket.Write(sendMsg)
		if err != nil {
			fmt.Println("fail to write udp msg, err : ", err)
			return
		}
	}
	time.Sleep(time.Hour)
}
