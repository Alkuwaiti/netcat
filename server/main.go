package main

import (
	"fmt"
	"net"
	"server/functions"
)

func main() {
	// Start listening on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Server started. Listening on :8080")

	initialMessage := functions.Welcome()

	// Accept incoming connections
	for {
		// this is a blocking line of code
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		fmt.Println("Client connected:", conn.RemoteAddr())

		// Handle connections in a new goroutine
		go handleConnection(conn, initialMessage)
	}
}

func handleConnection(conn net.Conn, initialMessage string) {
	defer conn.Close()

	_, err := conn.Write([]byte(initialMessage))
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err.Error())
		return
	}
	fmt.Print("this is the name: " + string(buffer))

	// infinite loop for the rest of the connection
	// for {

	// }

}
