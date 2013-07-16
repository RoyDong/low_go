package main

import (
    "low/app"
    "low/server"
)

func main(){ 
    app.LoadConfig()
    server.Start(app.Config.Port)
}
