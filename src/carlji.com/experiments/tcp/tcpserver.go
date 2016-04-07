package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9909")
	if err != nil {
		fmt.Printf("Unable resolve TCP address %v \n", err)
		return
	}
	listen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Errorf("Unable to list tcp %v \n", err)
		return
	}

	defer listen.Close()

	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Errorf("unabel to accept connection %v", err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn *net.TCPConn) {

	remoteAddr := conn.RemoteAddr().String()
	defer func() {
		fmt.Printf("Disconnected connection %v", remoteAddr)
		conn.Close()
	}()

	reader := bufio.NewReader(conn)

	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Errorf("Error read string %v", err)
		return
	}

	fmt.Println(string(message))
	conn.Write([]byte(time.Now().String() + "\n"))
	return
}
