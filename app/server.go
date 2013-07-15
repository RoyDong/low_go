package app

import (
    "log"
    "fmt"
    "net"
    "time"
    "bytes"
    "strings"
    "encoding/binary"
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

        if e == nil {
            go handle(conn)
        }else{
            log.Println(e)
        }
    }

    log.Println("Server shutdown")
}

func handle(conn net.Conn) {
    log.Println(fmt.Sprintf("Client %s(%s) connected",
            conn.RemoteAddr().Network(),
            conn.RemoteAddr().String()))
    defer conn.Close()
    defer log.Println(fmt.Sprintf("Client %s(%s) disconnected",
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
            log.Println("Expired command")
            continue
        }

        length, _ = binary.ReadVarint(bytes.NewBuffer(head[8:15]))
        stream := make([]byte, length)
        _, e = conn.Read(stream)

        if e != nil {
            log.Println(e)
            continue
        }

        command := strings.Split(string(stream), ".")

        if len(command) < 2 {
            log.Println("Error command:" + command)
            continue
        }

        length, _ = binary.ReadVarint(bytes.NewBuffer(head[16:24]))
        stream = make([]byte, length)
        _, e = conn.Read(stream)

        if e != nil {
            log.Println(e)
            continue
        }

        exec(command, stream)
    }
}
