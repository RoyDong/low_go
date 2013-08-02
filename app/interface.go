package app

import "net"

type Session interface {
    Close()
    Open()
    Conn() net.Conn
    Notify(data interface{})
    User() (User, bool)
    SetUser(u User)
}

type Message interface {
    GetInt(k string) (int64, bool)
    GetFloat(k string) (float64, bool)
    Get(k string) (string, bool)
    Id() byte
    Title() []byte
    TitlePath() []string
    Content() []byte
    Reply() []byte
    Replied() bool
    Session() Session
    SentAt() int64

    ReplySuccess(data interface{})
    ReplyError(code int64, message string)
    Send()
    SetReply(r []byte)
}

type User interface {
    Id() int64
    Name() string
    Session() (Session bool)
    SetSession(s Session)
    Json() []byte
}
