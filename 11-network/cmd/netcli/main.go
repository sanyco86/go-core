package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

const network, address = "tcp4", "localhost:8000"

func main() {
	fmt.Println("Подключаемся...")
	conn, err := net.Dial(network, address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	r := bufio.NewReader(os.Stdin)

	fmt.Println("Введите запрос:")

	for {
		query, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = conn.Write([]byte(query))
		if err != nil {
			fmt.Println(err)
			return
		}

		msg, err := io.ReadAll(conn)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(msg)
	}
}
