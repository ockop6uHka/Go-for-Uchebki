package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

type ConnectionHandler struct {
	Connection net.Conn
}

func NewConnectionHandler(conn net.Conn) *ConnectionHandler {
	return &ConnectionHandler{
		Connection: conn,
	}
}

func (ch *ConnectionHandler) Handle() {
	defer ch.Connection.Close()

	var wg sync.WaitGroup
	input := bufio.NewScanner(ch.Connection)
	for input.Scan() {
		wg.Add(1)
		go ch.echo(input.Text(), 1*time.Second, &wg)
	}
	wg.Wait()
}

func (ch *ConnectionHandler) echo(shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(ch.Connection, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(ch.Connection, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(ch.Connection, "\t", strings.ToLower(shout))
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		connHandler := NewConnectionHandler(conn)
		go connHandler.Handle()
	}
}
