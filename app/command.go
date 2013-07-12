package app

import (
    "log"
)

func exec(command string, data []byte) {
    log.Println(command)
    log.Println(string(data))
    path := strings.Split(command, ".")

    switch path[0] {
    case "user":
        switch path[1] {
        case "show": user.Show(data)
        case "login":

        default:
            log.Println("Command not found")
        }
    default:
        log.Println("Command not found")
    }
}
