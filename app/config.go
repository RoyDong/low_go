package app

import (
    "log"
    "os"
    "encoding/json"
)

const (
    ConfigFile = "config/app.json"
)

/*
    config data structure
*/
type mysql struct {
    Dsn string `json:"dsn"`
    Dbname string `json:"dbname"`
}

type config struct {
    Name string `json:"name"`
    Version string `json:"version"`
    Port int `json:"port"`
    Mysql mysql `json:"mysql"`
}

/*
    Config stores all the config data
*/
var Config = new(config)

func LoadConfig() {
    file, e := os.Open(ConfigFile)
    defer file.Close()

    if e != nil {
        log.Fatal(e)
    }

    fileInfo, e := file.Stat()

    if e != nil {
        log.Fatal(e)
    }

    stream := make([]byte, fileInfo.Size())
    file.Read(stream)
    e = json.Unmarshal(stream, Config)

    if e != nil {
        log.Fatal(e)
    }
}
