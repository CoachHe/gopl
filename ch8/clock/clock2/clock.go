package main

import (
	"io"
	"log"
	"net"
	"time"
)

//@author: coachhe
//@create: 2022/8/9 20:00

func main() {
	listener, err := net.Listen("tcp", "localhost:8000") // 监听本地的8080端口
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // 例如连接终止
			continue
		}
		go handleConn(conn) // 一次处理一个连接
	}
}

// 在这里，net.Conn满足io.Writer接口，因此可以直接通过io.WriteString()方法往里面写数据
func handleConn(c net.Conn) {
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			log.Print(err)
		}
	}(c)
	for {
		_, err := io.WriteString(c, time.Now().Format("15:05:04\n"))
		if err != nil {
			return // 例如连接断开
		}
		time.Sleep(1 * time.Second)
	}
}
