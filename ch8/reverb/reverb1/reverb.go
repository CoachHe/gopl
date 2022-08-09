package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

//@author: coachhe
//@create: 2022/8/9 20:35

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Println(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Println(c, "\t", shout)
	time.Sleep(delay)
	fmt.Println(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	// 注意： 忽略input.Err()中可能的错误
	err := c.Close()
	if err != nil {
		return
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Printf("conect failed, %v", err.Error())
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}
