package conn

import (
	"fmt"
	"io"
	"net"
	"os"
)

func Init() {
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
	rcv_loop(conn)

}

func rcv_loop(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from client: ", err.Error())
			os.Exit(1)
		}
	}
	conn.Write([]byte("Live\r\n"))
}
