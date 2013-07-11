package app

import (
    "log"
    "os"
    "encoding/json"
)

const (
    ConfigPath = "config/"
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
    Mysql mysql `json:"mysql"`
}

/*
    Config stores all the config data
*/
var Config = new(config)

func LoadConfig(name string) {
    file, e := os.Open(ConfigPath + name + ".json")

    if e != nil {
        log.Fatal(e)
        return
    }

    fileInfo, e := file.Stat()

    if e != nil {
        log.Fatal(e)
        return
    }

    stream := make([]byte, fileInfo.Size())
    file.Read(stream)
    e = json.Unmarshal(stream, Config)

    if e != nil {
        log.Fatal(e)
    }
}
