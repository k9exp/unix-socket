package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

func server(c net.Conn) {
    for {
        buf := make([]byte, 512)
        n, err := c.Read(buf)
        if err != nil {
            log.Fatal(err)
        }

        data := buf[0:n] // is a number
        num, err := strconv.Atoi(string(data))
        if err != nil {
            panic(err)
        }

        fmt.Println("Client Send: ", num)

        _, err = c.Write([]byte(fmt.Sprint(num + 1)))
        if err != nil {
            panic(err)
        }
    }
}

func serverRunner() {
    ln, err := net.Listen("unix", "/var/run/demo_2.sock")
    if err != nil {
        log.Fatal(err)
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Fatal(err)
        }

        go server(conn)
    }
}

func main(){
    serverRunner()
}

