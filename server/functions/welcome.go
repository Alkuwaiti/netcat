package functions

import (
	"bufio"
	"fmt"
	"os"
)

func Welcome() string {
	file, err := os.Open("./welcomeText.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer file.Close()

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	var linesFromFile string

	newLineStopper := 0

	// Iterate over each line and add it to the array
	for scanner.Scan() {
		newLineStopper += 1
		if newLineStopper != 18 {
			linesFromFile += scanner.Text() + "\n"
		} else {
			linesFromFile += scanner.Text()
		}
	}

	return linesFromFile

}
