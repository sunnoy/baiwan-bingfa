/*
 *@Description
 *@author          lirui
 *@create          2021-04-13 14:30
 */
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var RemoteAddr *string
var ConcurNum *int
var LocalAddr *string

func init() {
	RemoteAddr = flag.String("remote-ip", "127.0.0.1", "ip addr of remote server")
	ConcurNum = flag.Int("concurrent-num", 100, "concurrent number of client")
	LocalAddr = flag.String("local-ip", "0.0.0.0", "ip addr of remote server")
}

func consume() {

	laddr := &net.TCPAddr{IP: net.ParseIP(*LocalAddr)}

	var dialer net.Dialer
	dialer.LocalAddr = laddr

	conn, err := dialer.Dial("tcp", *RemoteAddr+":8888")
	if err != nil {
		fmt.Println("dial failed:", err)
		os.Exit(1)
	}
	defer conn.Close()

	buffer := make([]byte, 512)

	for {
		_, err2 := conn.Read(buffer)
		if err2 != nil {
			fmt.Println("Read failed:", err2)
			return
		}

		//  fmt.Println("count:", n, "msg:", string(buffer))

	}

}

func main() {
	flag.Parse()
	for i := 0; i < *ConcurNum; i++ {
		go consume()
	}
	time.Sleep(3600 * time.Second)
}