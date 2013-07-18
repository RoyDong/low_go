package server

import (
    "log"
    "fmt"
    "net"
    "time"
    "bytes"
    "encoding/binary"
    "encoding/json"
    "low/app"
)

const (
    StateClose = 0
    StateOpen = 1
    NoticeMid = byte(63)
)

type Session struct {
    listener net.Listener
    conn net.Conn
    user app.User
    state int
    openAt time.Time
}

type DataForJson struct {
    Code int64 `json:"code"`
    Message string `json:"message"`
    Data interface{} `json:"data"`
}

func (s *Session) Close() {
    s.conn.Close()
    s.state = StateClose
    log.Println(fmt.Sprintf("Client %s(%s) disconnected",
            s.conn.RemoteAddr().Network(),
            s.conn.RemoteAddr().String()))
}

func (s *Session) Open() {
    conn := s.conn
    s.state = StateOpen
    log.Println(fmt.Sprintf("Client %s(%s) connected",
            conn.RemoteAddr().Network(),
            conn.RemoteAddr().String()))
    defer s.Close()

    for {
        var length int32
        head := make([]byte, 13)
        _, e := conn.Read(head)

        if e != nil {
            log.Println(e)
            return
        }

        m := &Message{}
        m.session = s
        m.id = head[0]
        binary.Read(bytes.NewBuffer(head[1:5]), binary.LittleEndian, &m.sentAt)
        binary.Read(bytes.NewBuffer(head[5:9]), binary.LittleEndian, &length)

        if length == 0 {
            continue
        }

        m.title = make([]byte, length)
        _, e = conn.Read(m.title)

        if e != nil {
            log.Println(e)
            return
        }

        binary.Read(bytes.NewBuffer(head[9:13]), binary.LittleEndian, &length)

        if length == 0 {
            continue
        }

        m.content = make([]byte, length)
        _, e = conn.Read(m.content)

        if e != nil {
            log.Println(e)
            return
        }

        if netlag := time.Now().Unix() - int64(m.sentAt); netlag > 30 {
            log.Println("Expired command")
            continue
        }

        log.Println(m.id, m.sentAt, string(m.title), string(m.content))

        m.params = app.ParseHttpQuery(string(m.content))
        m.Dispatch()
        s.Reply(m)
    }
}

func (s *Session) Reply(m app.Message) {
    buffer := new(bytes.Buffer)
    binary.Write(buffer, binary.LittleEndian, m.Id())
    binary.Write(buffer, binary.LittleEndian, int32(time.Now().Unix()))
    binary.Write(buffer, binary.LittleEndian, int32(len(m.Reply())))
    binary.Write(buffer, binary.LittleEndian, m.Reply())

    s.conn.Write(buffer.Bytes())
}

func (s *Session) User() (app.User, bool) {
    if s.user == nil {
        return nil, false
    }

    return s.user, true
}

func (s *Session) SetUser(u app.User) {
    s.user = u
}

func (s *Session) Notify(data interface{}) {
    d, _ := json.Marshal(&DataForJson{
        Code: 0, Message: "done", Data: data,
    })

    buffer := new(bytes.Buffer)
    binary.Write(buffer, binary.LittleEndian, NoticeMid)
    binary.Write(buffer, binary.LittleEndian, int32(time.Now().Unix()))
    binary.Write(buffer, binary.LittleEndian, int32(len(d)))
    binary.Write(buffer, binary.LittleEndian, d)

    s.conn.Write(buffer.Bytes())
}
