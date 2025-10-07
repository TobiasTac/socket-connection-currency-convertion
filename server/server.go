package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

const exchangeRate = 0.20
const richUSD = 500.00

func classifyValue(valueUSD float64) string {
	if valueUSD >= richUSD {
		return "You Are Rich"
	} else {
		return "You Are Poor"
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Client Connect:", conn.RemoteAddr())
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Client disconnected or read error:", err)
			return
		}

		fmt.Printf("Message Received (BRL): %s", msg)

		trimmed := strings.TrimSpace(msg)

		valueBRL, err := strconv.ParseFloat(trimmed, 64)

		if err != nil || valueBRL < 0 {
			errMsg := "Error: You typed wrong\n"
			conn.Write([]byte(errMsg))
			continue
		}

		valueUSD := valueBRL * exchangeRate

		classification := classifyValue(valueUSD)

		response := fmt.Sprintf("BRL: R$%.2f. Converted to USD: $%.2f. Status: %s\n", valueBRL, valueUSD, classification)

		conn.Write([]byte(response))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:3000")

	if err != nil {
		fmt.Println("Error starting server", err)
		return
	}

	fmt.Println("🔥 Server started at localhost:3000")
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection", err)
			continue
		}

		go handleConnection(conn)
	}
}
