package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

	"example.com/personal_projects/redis/src/redis"
)

var queue = make(chan redis.Command, 100)

func main() {

	go redis.EventLoop(queue)

	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis server listening on :6379")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {

		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("client disconnected:", conn.RemoteAddr())
				return
			}
			fmt.Println("read error:", err)
			return
		}

		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		parts := strings.Fields(line)

		command := strings.ToUpper(parts[0])

		args := []string{}

		if len(parts) > 1 {
			args = parts[1:]
		}

		cmd := redis.Command{
			Name:   command,
			Args:   args,
			Result: make(chan redis.Response, 1),
		}

		queue <- cmd

		res := <-cmd.Result

		if res.Err != nil {
			conn.Write([]byte("ERR " + res.Err.Error() + "\n"))
			continue
		}

		conn.Write([]byte(fmt.Sprintf("%v\n", res.Data)))
	}
}
