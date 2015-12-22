package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"
)

const (
	MAX_CACHES = 1024
	TIMEOUT    = 30
)

func main() {
	var host string
	var port int
	args := os.Args[1:]
	argslen := len(args)
	fmt.Println(argslen)
	host = "127.0.0.1"
	port = 53
	if argslen >= 2 {
		host = args[0]
		newport, err := strconv.Atoi(args[1])
		if err != nil {
			port = 53
		} else {
			port = newport
		}
	}

	server := &Server{
		host:     host,
		port:     port,
		rTimeout: TIMEOUT * time.Second,
		wTimeout: TIMEOUT * time.Second,
	}

	server.Run()
	fmt.Println("DNS server start")

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)

forever:
	for {
		select {
		case <-sig:
			fmt.Println("Signal recieved, now stop and exit")
			break forever
		}
	}

}
