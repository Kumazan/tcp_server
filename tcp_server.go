package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

const (
	connType   = "tcp"
	connHost   = "localhost"
	connPort   = "2019"
	timeoutSec = 30
)

func main() {
	listener, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Listen error: ", err)
		return
	}

	fmt.Println("Starting listening...")
	for {
		c, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error: ", err)
			break
		}
		// Start a new goroutine to handle the new connection.
		go handleConn(c)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Please enter a distict in Taipei City (for example, 信義區, 松山區, etc.) to get the weather information or ‘quit’ to quit.\n"))

	scanner := bufio.NewScanner(conn)
	timeoutDuration := timeoutSec * time.Second
	for {
		conn.SetDeadline(time.Now().Add(timeoutDuration))
		if !scanner.Scan() || scanner.Text() == "quit" {
			break
		}
		conn.Write([]byte("Message received.\n"))

		keyword := scanner.Text()
		fmt.Println("Received keyword:", keyword, "from", conn.RemoteAddr())

		go requestExternalAPI(conn, keyword)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}
}
