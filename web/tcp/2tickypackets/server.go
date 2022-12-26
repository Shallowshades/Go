/*
TCP Server 黏包 案例
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		_, err := reader.Read(buf[:])
		if err == io.EOF {
			fmt.Println("clent break link...")
			break
		}
		if err != nil {
			fmt.Println("faid to read client, err : ", err)
			break
		}
		recvStr := string(buf[:])
		fmt.Println("CLIENT : ", recvStr)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("fail to listen, err : ", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("failed to accept, err : ", err)
			continue
		}
		go process(conn)
	}
}
