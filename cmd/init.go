package main

import (
	"fmt"
	"net"
)

func init() {
	lstn, err := net.Listen("tcp", ":6889")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := lstn.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

}
