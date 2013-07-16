package user


import (
    "log"
    "low/server"
)

func Show(m *server.Message) {
    log.Println(m.Command())
}
