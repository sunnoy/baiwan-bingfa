/*
 *@Description
 *@author          lirui
 *@create          2021-04-13 14:30
 */
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var array []byte = make([]byte, 10)


func Handler(conn net.Conn) {
	for {
		_, err := conn.Write(array)
		if err != nil {
			return
		}
		time.Sleep(20 * time.Second)
	}
}

func main() {

	for i := 0; i < 10; i += 1 {
		array[i] = 'a'
	}

	service := ":8888"
	// 构造tcp端点
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)
	// 监听
	l, _ := net.ListenTCP("tcp", tcpAddr)

	for {
		// accept
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("accept error, err=%s\n", err.Error())
			os.Exit(1)
		}
		go Handler(conn)
	}

}
