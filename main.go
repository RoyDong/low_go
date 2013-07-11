package main

import (
    "log"
    "low/app"
    "low/user"
)

func main(){
    app.LoadConfig("app")

    u, _ := user.Find(7)

    log.Println(u.Json())

    log.Println("_______")
}
