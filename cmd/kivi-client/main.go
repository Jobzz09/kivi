package main

import (
	"fmt"

	"github.com/Jobzz09/kivi/internal/conn"
)

func main() {
	conn.Init()
	fmt.Println("Started client")
}
