package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := Decode(reader)
		if err == io.EOF {
			fmt.Println("conn.Read() err : ", err)
			return
		}
		if err != nil {
			fmt.Println("failed to decode msg, err : ", err)
			return
		}
		fmt.Println("CLIENT : ", msg)

		data, err := Encode(msg)
		if err != nil {
			fmt.Println("fail to encode msg, err : ", err)
			return
		}
		conn.Write(data)
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
			fmt.Println("fail to accept, err : ", err)
			continue
		}
		go process(conn)
	}
}

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
	//读取消息的长度，转换成int类型
	length := int32(len(message))
	pkg := new(bytes.Buffer)
	//写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	//写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	//读取消息的长度，转换成int
	var length int32
	lengthByte, _ := reader.Peek(4)                              //读取前四个字节的数据 []byte
	lengthBuff := bytes.NewBuffer(lengthByte)                    //*bytes.Buffer
	err := binary.Read(lengthBuff, binary.LittleEndian, &length) //int32
	if err != nil {
		return "", err
	}
	//Buffered返回缓冲中现有的可读取的字节数
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	//读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
