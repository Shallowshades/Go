package main

import (
	"fmt"
	"net"
)

// func Write(msg []byte) {

// }

// func Read(reader bufio.Reader) {

// }

func main() {

	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("fail to listen, err : ", err)
		return
	}
	defer listen.Close()
	for {
		msg := make([]byte, 512) //没有内存的情况，下面将无法从udp链接中读取数据
		n, addr, err := listen.ReadFromUDP(msg)
		if err != nil {
			fmt.Println("fail to read udp msg, err : ", err)
			continue
		}
		fmt.Printf("addr[%v], count(%v), msg : %v\n", addr, n, string(msg[:n]))
		_, err = listen.WriteToUDP(msg[:n], addr)
		if err != nil {
			fmt.Println("fail to write udp msg, err : ", err)
			continue
		}
	}
}
