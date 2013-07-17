package app

type Session interface {
    Close()
    Open()
    Reply(m Message)
}

type Message interface {
    GetInt(k string) (int64, bool)
    Get(k string) (string, bool)
    Cid() byte
    Title() []byte
    TitlePath() []string
    Content() []byte
    Reply() []byte
    Replied() bool
    Session() Session
    SentAt() int32
    SetReply(data []byte)
    Read(content interface{})
}
