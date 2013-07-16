package server

import (
    "log"
    "fmt"
    "net"
    "time"
    "bytes"
    "encoding/binary"
    "encoding/json"
)

type Session struct {
    listener net.Listener
    conn net.Conn
    startTime int64
}

func Start(port int) {
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
                startTime: time.Now().Unix(),
            }

            go session.Open()
        }else{
            log.Println(e)
        }
    }

    log.Println("Server shutdown")
}

func (s *Session)Close() {
    s.conn.Close()
    log.Println(fmt.Sprintf("Client %s(%s) disconnected",
            s.conn.RemoteAddr().Network(),
            s.conn.RemoteAddr().String()))
}

func (s *Session)Open() {
    conn := s.conn
    log.Println(fmt.Sprintf("Client %s(%s) connected",
            conn.RemoteAddr().Network(),
            conn.RemoteAddr().String()))
    defer s.Close()

    for {
        var callTime, length int32
        head := make([]byte, 13)
        _, e := conn.Read(head)

        if e != nil {
            log.Println(e)
            return
        }

        m := &Message{}
        m.id = head[0]
        binary.Read(bytes.NewBuffer(head[1:5]), binary.LittleEndian, &callTime)

        binary.Read(bytes.NewBuffer(head[5:9]), binary.LittleEndian, &length)
        m.command = make([]byte, length)
        _, e = conn.Read(m.command)

        if e != nil {
            log.Println(e)
            return
        }

        binary.Read(bytes.NewBuffer(head[9:13]), binary.LittleEndian, &length)
        m.data = make([]byte, length)
        _, e = conn.Read(m.data)

        if e != nil {
            log.Println(e)
            return
        }

        if netlag := time.Now().Unix() - int64(callTime); netlag > 30 {
            log.Println("Expired command")
            continue
        }

        m.Handle()
        m.Response()
    }
}

type Message struct {
    id byte
    command, data, body []byte
    sendTime int32
    session *Session
}

func (m *Message)Cid() byte {
    return m.id
}

func (m *Message)Command() string {
    return string(m.command)
}

func (m *Message)Read(data interface{}) {
    json.Unmarshal(m.data, data)
}

func (m *Message)Session() *Session {
    return m.session
}

func (m *Message)SetBody(data []byte) {
    m.body = data
}

func (m *Message)Response() {
    t := byte(time.Now().Unix())
    length := byte(len(m.body))
    head := make([]byte, 9 + len(m.body))

    head[0] = m.id
    head[1] = t & 7
    head[2] = t >> 8 & 7
    head[3] = t >> 16 & 7
    head[4] = t >> 24
    head[5] = length & 7
    head[6] = length >> 8 & 7
    head[7] = length >> 16 & 7
    head[8] = length >> 24

    copy(head, m.body)
    m.session.conn.Write(head)
}
