package server

import (
    "log"
    "low/controller/user"
)

func (m *Message)Dispatch() {
    path := m.TitlePath()

    switch path[0] {
    case "user":

        switch path[1] {
        case "show":user.Show(m)

        default:
            log.Println("no action")
        }

    default:
        log.Println("no action")
    }
}
