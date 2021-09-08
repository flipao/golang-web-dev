package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type request struct {
	method, url, body string
	headers           map[string]string
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	r := read(conn)

	fmt.Println(r)
}

func read(conn net.Conn) request {
	var m, u, b string
	h := make(map[string]string)
	i := 0
	doneHeaders := false

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
		if i == 0 {
			l := strings.Fields(line)
			m = l[0]
			u = l[1]
		}
		if !doneHeaders {
			if line == "" {
				doneHeaders = true
				break
			} else {
				l := strings.Fields(line)
				h[l[0]] = l[1]
			}
		} else {
			b += line
		}

		i++
	}

	return request{
		method:  m,
		url:     u,
		body:    b,
		headers: h,
	}
}
