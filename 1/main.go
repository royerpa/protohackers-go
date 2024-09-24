package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
)

type Request struct {
	Method *string  `json:"method"`
	Number *float64 `json:"number"`
}

type Response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

type Error struct {
	Error string `json:"error"`
}

func isMalformedRequest(req *Request) bool {
	return req.Method == nil || *req.Method != "isPrime" || req.Number == nil
}

func isPrime(n float64) bool {
	if n != float64(int(n)) || n < 2 {
		return false
	}

	num := int(n)
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}

	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	encoder := json.NewEncoder(conn)

	for scanner.Scan() {
		var req Request
		if err := json.Unmarshal(scanner.Bytes(), &req); err != nil || isMalformedRequest(&req) {
			encoder.Encode(Error{Error: "malformed request"})
			continue
		}

		encoder.Encode(Response{Method: "isPrime", Prime: isPrime(*req.Number)})
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from connection: %v", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":7")
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer listener.Close()

	log.Println("Prime server is listening on port 7")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		log.Printf("Accepted connection from %v", conn.RemoteAddr())
		go handleConnection(conn)
	}
}
