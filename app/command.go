package app

import (
    "log"
    "controller/user"
)

func exec(command []string, data []byte) {

    switch command[0] {
    case "user":
        switch command[1] {
        case "show": user.ShowAction(data)
        case "login":

        default:
            log.Println("Command not found")
        }
    default:
        log.Println("Command not found")
    }
}
