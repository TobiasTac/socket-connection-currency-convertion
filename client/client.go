package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("Error: Server Not Found", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connecting server...")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Enter the amount in R$: [type 'exit' to exit]\n")
		text, _ := reader.ReadString('\n')

		if strings.Contains(strings.ToLower(text), "exit") {
			fmt.Print("Connection Terminated.\n")
			break
		}

		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error sending data", err)
			break
		}

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading server response", err)
			break
		}
		fmt.Print("Response: " + message + "\n")
	}
}
