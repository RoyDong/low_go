package app

type Session interface {
    Close()
    Open()
    Reply(m Message)
    Notify(data interface{})
    User() (User, bool)
    SetUser(u User)
}

type Message interface {
    GetInt(k string) (int64, bool)
    Get(k string) (string, bool)
    Id() byte
    Title() []byte
    TitlePath() []string
    Content() []byte
    Reply() []byte
    Replied() bool
    Session() Session
    SentAt() int32

    ReplySuccess(data interface{})
    ReplyError(code int64, message string)
    SetReply(r []byte)
}

type User interface {
    Id() int64
    Name() string
    Json() []byte
}
