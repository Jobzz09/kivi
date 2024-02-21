package main

import (
	"fmt"
	"time"

	"github.com/Jobzz09/kivi/internal/conn"
)

func main() {
	go conn.Init()
	fmt.Println("Started server")
	time.Sleep(time.Minute * 60)
}
