package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	// Create a new scanner reading from the standard input (os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)

	// Read response from the server
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err.Error())
		return
	}

	fmt.Print(string(buffer))

	// Scan for the next token (which by default is a line)
	scanner.Scan()

	// Retrieve the user's name that they entered
	name := scanner.Text()

	_, err = conn.Write([]byte(name))
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	// retrieve all chat log before sending a new message

	fmt.Print("Enter a message: ")

	scanner.Scan()

	// Retrieve the user's name that they entered
	userInput := scanner.Text()

	_, err = conn.Write([]byte(userInput))
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	// for {

	// }

}
