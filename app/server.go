package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	// TEST command: echo -ne "" | nc localhost 6379

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)

		if err != nil {
			// if reach the EOF of input of request, close connection
			if err.Error() == "EOF" {
				fmt.Println("Connection closed")
				conn.Close()
				return
			}
			fmt.Println("Error reading data from conn: ", err.Error())
			os.Exit(1)
		}

		fmt.Println(string(buf))
		message := []byte("+PONG\r\n")
		conn.Write(message)
	}
}
