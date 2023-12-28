package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Client struct {
	ServerAddr string
}

func NewClient(serverAddr string) *Client {
	return &Client{
		ServerAddr: serverAddr,
	}
}

func (c *Client) Run() {
	serverTCPAddr, err := net.ResolveTCPAddr("tcp", c.ServerAddr)
	if err != nil {
		log.Fatal(err)
	}

	serverConn, err := net.DialTCP("tcp", nil, serverTCPAddr)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		c.copyDataFromServer(serverConn)
		log.Println("data transfer completed")
		done <- struct{}{}
	}()

	go func() {
		c.copyDataToServer(serverConn, os.Stdin)
		log.Println("client exiting")
		serverConn.CloseWrite()
	}()

	// Handle Ctrl+C or termination signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	serverConn.Close()
	<-done
}

func (c *Client) copyDataToServer(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}
}

func (c *Client) copyDataFromServer(src io.Reader) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if _, err := io.Copy(os.Stdout, src); err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}
	}()

	wg.Wait()
}

func main() {
	serverAddr := "localhost:8000"
	client := NewClient(serverAddr)
	client.Run()
}
