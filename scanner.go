package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

const Timeout = 1 * time.Second

func scan (host string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)

	connection, err := net.DialTimeout("tcp", address, Timeout)

	if err != nil{
		return
	}
	
	fmt.Println("Remote address:", connection.RemoteAddr().String())
	fmt.Printf("\tport %d is open\n", port)

	connection.Close()
}

func main() {

	var host string
	var port int

	flag.StringVar(&host, "host", "", "ip or root domain")
	flag.IntVar(&port, "port", 80, "port")

	flag.Parse()
	
	scan(host, port)
}
