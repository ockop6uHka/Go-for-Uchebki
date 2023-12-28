package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listenAddress := "localhost:8000"
	server, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	messageChannel := make(chan string)
	defer c.Close()

	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			messageChannel <- input.Text()
		}
		close(messageChannel)
	}()

	for {
		select {
		case message, ok := <-messageChannel:
			if !ok {
				return
			}
			go processMessage(c, message, 1*time.Second)
		case <-time.After(10 * time.Second):
			return
		}
	}
}

func processMessage(c net.Conn, message string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(message))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", message)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(message))
}
