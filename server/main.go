package main

import (
	"fmt"
	"net"
	"server/functions"
	"time"
)

type Data struct {
	Name    string
	Date    time.Time
	Message string
}

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

	messagesChannel := make(chan Data)

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
		go handleConnection(conn, initialMessage, messagesChannel)
	}
}

func handleConnection(conn net.Conn, initialMessage string, messagesChannel chan Data) {

	// to handle closing connection
	defer conn.Close()

	// Send the client the initial message
	_, err := conn.Write([]byte(initialMessage))
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	// define a buffer to re-use every now and then
	buffer := make([]byte, 1024)

	// read client's name
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err.Error())
		return
	}

	// save the name in a variable
	name := string(buffer)
	fmt.Println("this is the name: " + name)

	// write to the client all previous messages before infinite loop ---------- right here

	// Read client's message
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err.Error())
		return
	}

	// save user's message in a variable
	message := string(buffer)

	currentTime := time.Now()

	// this is a blocking line
	messagesChannel <- Data{Name: name, Date: time.Now(), Message: message}

	close(messagesChannel)

	fmt.Println("[" + currentTime.Format("2006-01-02 15:04:05") + "][" + name + "]:" + message)

	// infinite loop for the rest of the connection
	// for {

	// }

}
