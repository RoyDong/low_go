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
    m.ReplySuccess(u.Data())
}

func Signin(m app.Message) {
    email, has := m.Get("email")

    u, has := user.FindBy("email", email)

    if !has {
        m.ReplyError(2, "sss")
        return
    }

    m.Session().Notify(u.Data())

    m.ReplySuccess(u.Data())
}
