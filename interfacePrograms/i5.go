package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() { // Scans a line from Stdin(Console)
		//strLine := strings.TrimSpace(scanner.Text())
		fmt.Println(">>>>>>")
		strLine := scanner.Text() // Holds the string that scanned
		fmt.Println("strLine:" + strLine + ":")
	}

}
