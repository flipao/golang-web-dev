package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
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

	// read request
	u := request(conn)

	// write response
	respond(conn, u)
}

func request(conn net.Conn) string {
	i := 0
	var url string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			r := strings.Fields(ln)
			fmt.Println("***METHOD", r[0])
			url = r[1]
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	return url
}

func respond(conn net.Conn, u string) {

	body := fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World %s</strong></body></html>`, u)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
