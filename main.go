package main

import (
    "log"
    "low/app"
    "low/user"
)

func main(){
    app.LoadConfig()

    u, _ := user.Find(7)

    user.Delete(u.Id())


    log.Println("_______")
}
