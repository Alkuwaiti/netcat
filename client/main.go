package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	for {
		// Send a message to the server
		message := "Hello, server!"
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending:", err.Error())
			return
		}
		fmt.Println("Message sent:", message)

		// Read response from the server
		buffer := make([]byte, 1024)
		_, err = conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading response:", err.Error())
			return
		}
		fmt.Println("Response from server: " + string(buffer))
	}

}
