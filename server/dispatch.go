package server

import (
    "log"
    "strings"
    //"low/controller/user"
)

func (m *Message)Handle() {
    path := strings.Split(string(m.command), ".")

    log.Println(path[0])

    /*
    switch path[0] {
    case "user":

        switch path[1] {
        case "show": user.Show(m)
        case "login":

        default:
            log.Println("Command not found")
        }
    default:
        log.Println("Command not found")
    }
    */
}
