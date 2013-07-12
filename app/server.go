package app

import (
    "log"
    "fmt"
    "net"
    "time"
    "encoding/binary"
    "bytes"
)

func Run() {
    LoadConfig()
    listener, e := net.Listen("tcp", fmt.Sprintf(":%d", Config.Port))
    defer listener.Close()
    log.Println(fmt.Sprintf("Server started, listening port:%d", Config.Port))

    if e != nil {
        log.Fatal(e)
    }

    for {
        conn, e := listener.Accept()

        if e != nil {
            log.Println(e)
            continue
        }

        go handle(conn)
    }

    log.Println("Server shutdown")
}

func handle(conn net.Conn) {
    log.Println(fmt.Sprintf("Client: %s(%s)",
            conn.RemoteAddr().Network(),
            conn.RemoteAddr().String()))

    for {
        var length int64
        head := make([]byte, 24)
        _, e := conn.Read(head)

        if e != nil {
            log.Println(e)
            continue
        }

        startTime, _ := binary.ReadVarint(bytes.NewBuffer(head[0:7]))
        netlag := time.Now().Unix() - startTime

        log.Println(startTime)

        if netlag < 0 || netlag > 30 {
            log.Println("Command timeout")
            continue
        }

        length, _ = binary.ReadVarint(bytes.NewBuffer(head[8:15]))
        command := make([]byte, length)
        _, e = conn.Read(command)

        if e != nil {
            log.Println(e)
            continue
        }

        length, _ = binary.ReadVarint(bytes.NewBuffer(head[16:24]))
        body := make([]byte, length)
        _, e = conn.Read(body)

        if e != nil {
            log.Println(e)
            continue
        }

        exec(string(command), body)
    }
}
