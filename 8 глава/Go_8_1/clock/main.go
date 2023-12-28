package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ClockServer struct {
	Address  string
	TimeZone string
}

type ClockHandler interface {
	HandleClock(conn net.Conn)
}

type DefaultClockHandler struct {
	TimeZone string
	Output   io.Writer
}

func (d *DefaultClockHandler) HandleClock(conn net.Conn) {
	defer conn.Close()
	for {
		loc, err := time.LoadLocation(d.TimeZone)
		if err != nil {
			log.Fatal(err)
		}
		currentTime := time.Now().In(loc).Format("2006-01-02 15:04:05 MST")

		_, err = io.WriteString(conn, currentTime)
		if err != nil {
			return
		}

		_, err = io.WriteString(d.Output, currentTime+"\n")
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}

type ClockServerManager struct {
	Servers []ClockServer
	Handler ClockHandler
	wg      sync.WaitGroup
}

func (m *ClockServerManager) StartServers() {
	for _, server := range m.Servers {
		m.wg.Add(1)
		go func(server ClockServer) {
			defer m.wg.Done()
			conn, err := net.Dial("tcp", server.Address)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()

			m.Handler.HandleClock(conn)
		}(server)
	}
}

func (m *ClockServerManager) Wait() {
	m.wg.Wait()
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("Usage: clockwall <server1:port1:timezone1> <server2:port2:timezone2> ...")
	}

	var servers []ClockServer
	for _, arg := range args {
		parts := strings.Split(arg, ":")
		if len(parts) != 3 {
			log.Fatal("Неправильный формат аргументов. Используйте <server:port:timezone>")
		}
		port, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal("Неправильное число порта")
		}

		servers = append(servers, ClockServer{
			Address:  parts[0] + ":" + strconv.Itoa(port),
			TimeZone: parts[2],
		})
	}

	handler := &DefaultClockHandler{TimeZone: "UTC", Output: os.Stdout}
	manager := &ClockServerManager{Servers: servers, Handler: handler}

	// Добавлено для запуска локального сервера часов
	localTimeZone := "UTC"
	localPort := 8000
	localServer := ClockServer{Address: "localhost:" + strconv.Itoa(localPort), TimeZone: localTimeZone}
	localHandler := &DefaultClockHandler{TimeZone: localTimeZone, Output: os.Stdout}
	localManager := &ClockServerManager{Servers: []ClockServer{localServer}, Handler: localHandler}

	fmt.Printf("Ожидание соединений на порту %d...\n", localPort)

	go localManager.StartServers() // Запуск локального сервера в фоновом режиме
	manager.StartServers()
	manager.Wait()
}
