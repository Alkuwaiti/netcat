package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	// Read initial message from the server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	fmt.Println("Server:", string(buffer[:n]))

	// Read user's name from stdin
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Send the name to the server
	_, err = conn.Write([]byte(name))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
		return
	}
	currentTime := time.Now()

	// Goroutine to continuously read from the server
	go func() {
		for {
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading from server:", err.Error())
				return
			}
			fmt.Println()
			fmt.Println(message)
			fmt.Print("[" + currentTime.Format("2006-01-02 15:04:05") + "][" + name + "]:")
		}
	}()

	// Continuosly send messages to the server
	for {

		fmt.Print("[" + currentTime.Format("2006-01-02 15:04:05") + "][" + name + "]:")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if strings.ToLower(message) == "exit" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		// Send message to the server
		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}
	}
}
