package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // Listen on TCP port 8080
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer ln.Close()

    fmt.Println("TCP server listening on port 8080")

    for {
        // Accept a connection
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }

        // Handle the connection in a new goroutine
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    // Read request from the connection
	request := make([]byte, 200)
    n, err := bufio.NewReader(conn).Read(request) 
	//ReadString('\n')
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Received request of %v: %s", n,  request)

    // Respond with "Hello World"
    _, err = conn.Write([]byte("HTTP/1.1 200 OK\r\nDate: Mon, 20 Nov 2023 18:44:58 GMT\r\nConnection: keep-alive\r\nKeep-Alive: timeout=5\r\nContent-Length: 13\r\n\r\nhello worlddd"))
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Sent response: Hello World")
}
