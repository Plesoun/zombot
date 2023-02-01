package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "zombot/bot"
    "log"
    "os"
)

type configStruct struct {
    Token     string `json:"token"`
    botPrefix string `json:"bot_prefix"`
}

var (
    token     string
    botPrefix string
    config *configStruct
)

func readConfig() error {
    // TODO: change to log message later
    fmt.Println("Reading config file")
    file, err := ioutil.ReadFile("./config.json")
    if err != nil {
        fmt.Println("Error loading config file...", err.Error())
        return err
    }
    err = json.Unmarshal(file, &config)
    if err != nil {
        fmt.Println("Error parsing config file...", err.Error())
        return err
    }

    token = config.Token
    botPrefix = config.botPrefix

    return nil
}

func main() {
    // read config
    err := readConfig()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("Zomb!~")
}