package conn

import (
	"fmt"
	"net"
	"strings"

	"github.com/Jobzz09/kivi/internal/aof"
	"github.com/Jobzz09/kivi/internal/handler"
	"github.com/Jobzz09/kivi/internal/resp"
)

func Init() {
	lstn, err := net.Listen("tcp", ":6388")
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

	aof, err := aof.NewAof("database.aof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer aof.Close()

	//aof.Read(resp.Value{})

	rcv_loop(conn, aof)
}

func rcv_loop(conn net.Conn, _aof *aof.Aof) {
	for {
		_resp := resp.NewResp(conn)
		value, err := _resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}
		if value.Typ != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}

		if len(value.Array) == 0 {
			fmt.Println("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		writer := resp.NewWriter(conn)
		handler, ok := handler.Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(resp.Value{Typ: "string", Str: ""})
			continue
		}

		if command == "SET" || command == "HSET" {
			_aof.Write(value)
		}

		result := handler(args)
		writer.Write(result)
	}
}
