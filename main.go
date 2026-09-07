package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// 1. Listen for incoming TCP connections on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	fmt.Println("Server running on http://localhost:8080")

	for {
		// 2. Accept incoming connection
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		// Handle connection in a goroutine so it doesn't block other requests
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 3. Read the text file from disk
	content, err := os.ReadFile("sample.txt")
	if err != nil {
		// Send HTTP 404 response if file is missing or unreadable
		notFoundResponse := "HTTP/1.1 404 Not Found\r\n" +
			"Content-Type: text/plain; charset=utf-8\r\n" +
			"Content-Length: 14\r\n" +
			"Connection: close\r\n\r\n" +
			"File not found!"
		conn.Write([]byte(notFoundResponse))
		return
	}

	// 4. Construct raw HTTP/1.1 response
	responseHeader := fmt.Sprintf(
		"HTTP/1.1 200 OK\r\n"+
			"Content-Type: text/plain; charset=utf-8\r\n"+
			"Content-Length: %d\r\n"+
			"Connection: close\r\n\r\n",
		len(content),
	)

	// 5. Write HTTP headers and content back to the socket
	conn.Write([]byte(responseHeader))
	conn.Write(content)
}
