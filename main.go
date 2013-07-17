package main

import (
    "low/app"
    "low/server"
)

func main(){ 
    app.LoadConfig()
    server.Listen(app.Config.Port)
}
