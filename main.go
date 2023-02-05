package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "zombot/bot"
)

type configStruct struct {
    Token     string
    BotPrefix string
    DebugLogLocation string
}

var (
    config *configStruct
)

func readConfig() error {
    file, err := os.ReadFile("./config.json")
    if err != nil {
        log.Fatal("Error loading config file...", err.Error())
        return err
    }
    err = json.Unmarshal(file, &config)
    if err != nil {
        log.Fatal("Error parsing config file...", err.Error())
        return err
    }

    return nil
}

func main() {
    // read config
    err := readConfig()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("Zomb!", config)
    bot.DiscordBotToken = config.Token
    bot.Run()
}