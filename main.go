package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	message := map[string]string{
		"message": "Hello, ELK from Go!",
	}

	conn, err := net.Dial("tcp", "localhost:5044")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	msg, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(conn, "%s\n", msg)
	fmt.Println("Log sent:", string(msg))
}
