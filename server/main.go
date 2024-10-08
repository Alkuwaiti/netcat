package main

import (
	"fmt"
	"net"
	"server/functions"
	"time"
)

type Data struct {
	Name         string
	Date         time.Time
	Message      string
	LocalAddress net.Addr
}

// TODO :- Retrieve all messages before a client connects
// TODO :- Handle ^C cases where client leaves
// TODO :- Handle new joiners

var connectedClients []net.Conn

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

	// Goroutine to handle sending messages to clients
	go func() {
		for {
			// recieve messages
			data := <-messagesChannel
			// Send data to all connected clients
			for _, conn := range connectedClients {
				// if the Remote Address only isn't the same as Local Address (Won't send to the same client their own message)
				if conn.RemoteAddr() != data.LocalAddress {
					_, err := conn.Write([]byte(data.Message))
					if err != nil {
						fmt.Println("Error writing to client:", err.Error())
					}
				}
			}
		}
	}()

	// Accept incoming connections
	for i := 0; i < 10; i++ {
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

	// Send Client the previous messages here (in a different go routine? I think I need a waitGroup too)

	// save the name in a variable
	name := string(buffer)
	fmt.Println("this is the name: " + name)

	// Add client to connected clients list
	connectedClients = append(connectedClients, conn)

	// Handle receiving and sending messages
	for {
		// Read client's message
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading response:", err.Error())
			break
		}

		message := string(buffer[:n])

		currentTime := time.Now()

		fmt.Print("[" + currentTime.Format("2006-01-02 15:04:05") + "][" + name + "]:" + message)

		// Send message to messagesChannel to broadcast to all clients
		messagesChannel <- Data{Name: name, Date: currentTime, Message: "[" + currentTime.Format("2006-01-02 15:04:05") + "][" + name + "]:" + message, LocalAddress: conn.RemoteAddr()}
	}

}
