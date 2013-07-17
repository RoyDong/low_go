package user


import (
    "log"
    "low/app"
    "low/user"
)

func Show(m app.Message) {
    id, _ := m.GetInt("id")


    u, _ := user.Find(id)

    log.Println(id)
    m.SetReply(u.Json())
}
