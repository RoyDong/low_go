package server

import (
    "log"
    "fmt"
    "net"
    "time"
)

func Listen(port int) {
    listener, e := net.Listen("tcp", fmt.Sprintf(":%d", port))
    defer listener.Close()
    log.Println(fmt.Sprintf("Server started, listening port:%d", port))

    if e != nil {
        log.Fatal(e)
    }

    for {
        conn, e := listener.Accept()

        if e == nil {
            session := &Session{
                listener: listener,
                conn: conn,
                openAt: time.Now(),
            }

            go session.Open()
        }else{
            log.Println(e)
        }
    }

    log.Println("Server shutdown")
}
