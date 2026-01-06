package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"github.com/realaryann/keystore/resp"
	"github.com/realaryann/keystore/cache"
	"github.com/realaryann/keystore/command"
)


func main() {
	// Key: [value, timestamp, expiry]
	c := cache.Cache{Data: make(map[string][]string)}
	fmt.Println("KeyStore [S]")
	tcpl, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Infinite Loop to answer connections
	for {

		conn, err := tcpl.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
		}

		go HandleCon(conn, &c)

	}
}

func HandleCon(conn net.Conn, c *cache.Cache) {
	defer conn.Close()
	response := resp.NewReader(conn)	
	write := resp.NewWriter(conn)
	for {

		value, err := response.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error: ", err.Error())
			os.Exit(1)	

		}
		
		retval := command.Process(value, c)
		write.Write(retval)
	}

}
