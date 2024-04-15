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

	// Read response from the server
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err.Error())
		return
	}
	fmt.Println("Response from server: ")
	fmt.Println(string(buffer))

	// for {

	// }

}
