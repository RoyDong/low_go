package server

import (
    "strings"
    "low/app"
    "strconv"
    "encoding/json"
)

type Message struct {
    id byte
    title, content, reply []byte
    params map[string]string
    replied bool
    session app.Session
    sentAt int32
}

func (m *Message) GetInt(k string) (int64, bool) {
    if v, has := m.params[k]; has {
        if i, e := strconv.ParseInt(v, 10, 64); e == nil {
            return i, true
        }
    }

    return 0, false
}

func (m *Message) Get(k string) (string, bool) {
    if v, has := m.params[k]; has {
        return v, true
    }

    return "", false
}

func (m *Message) Id() byte {
    return m.id
}

func (m *Message) Title() []byte {
    return m.title
}

func (m *Message) TitlePath() []string {
    return strings.Split(string(m.title), ".")
}

func (m *Message) Content() []byte {
    return m.content
}

func (m *Message) Reply() []byte {
    return m.reply
}

func (m *Message) Replied() bool {
    return m.replied
}

func (m *Message) Session() app.Session {
    return m.session
}

func (m *Message) SentAt() int32 {
    return m.sentAt
}

func (m *Message) ReplySuccess(data interface{}) {
    m.reply, _ = json.Marshal(&DataForJson{
        Code: 0, Message: "done", Data: data,
    })
    m.replied = true
}

func (m *Message) ReplyError(code int64, message string) {
    m.reply, _ = json.Marshal(&DataForJson{
        Code: code, Message: message, Data: nil,
    })
    m.replied = true
}

func (m *Message) SetReply(r []byte) {
    m.reply = r
    m.replied = true
}
