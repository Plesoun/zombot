package bot

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
    "log"
    "os"
    "os/exec"
    "os/signal"
    "strings"
    "zombot/zomblogs"
)

var (
    DiscordBotToken string
    CommandPrefix string
    DebugLogLocation string
)

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
    // If it is a bot message, ignore
    if message.Author.ID == discord.State.User.ID {
        return
    }
    // if not, respond
    switch {
        // list all commands here TODO: maybe some better structure
        case strings.Contains(message.Content, CommandPrefix + "help"):
            discord.ChannelMessageSend(message.ChannelID, "Available commands: \n !help \n !status \n !logouts \n !logins \n !system \n !lasthorde")
        // system returns basic system stats from "top"
        case strings.Contains(message.Content, CommandPrefix +"system"):
            out, err := exec.Command("bash", "-c", "top -b -n 1 | egrep 'top -|Tasks:|%Cpu|MiB'").Output()
            if err != nil {
                log.Fatal("Encountered error while executing system command: ", err)
            }
            discord.ChannelMessageSend(message.ChannelID, string(out[:]))
            // status returns the systemd status section relevant to the query
        case strings.Contains(message.Content, CommandPrefix +"status"):
            out, err := exec.Command("bash", "-c", "sudo systemctl status zomboid.service | grep Active").Output()
            if err != nil {
                log.Fatal("Encountered error while executing system command: ", err)
            }
            discord.ChannelMessageSend(message.ChannelID, string(out[:]))
        // see all logouts contained in the current log file TODO: limit to ~ last 10
        case strings.Contains(message.Content, CommandPrefix + "logouts"):
            parsedLogs, err := parser.ParseLogFile(DebugLogLocation)
            if err != nil {
                log.Fatal("Error while parsing the log file:", err)
            }
            for _, logEntry := range parsedLogs {
                if strings.Contains(logEntry.Event, "disconnected") {
                    discord.ChannelMessageSend(message.ChannelID, fmt.Sprintf("%v %s %s", logEntry.Timestamp, logEntry.Name, logEntry.Event))
                }
            }
        // see all logins contained in the current log file TODO: limit to ~ last 10
        // TODO: file location and name should be configurable (rather directory containing logs should
        case strings.Contains(message.Content, CommandPrefix + "logins"):
            parsedLogs, err := parser.ParseLogFile(DebugLogLocation)
            if err != nil {
                log.Fatal("Error while parsing log file:", err)
            }
            fmt.Print(parsedLogs)
            for _, logEntry := range parsedLogs {
                fmt.Println(logEntry.Event)
                if strings.Contains(logEntry.Event, "fully connected") {
                    discord.ChannelMessageSend(message.ChannelID, fmt.Sprintf("%v %s %s", logEntry.Timestamp, logEntry.Name, logEntry.Event))
                }
            }
        case strings.Contains(message.Content, CommandPrefix +"lasthorde"):
            parsedLogs, err := parser.ParseLogFile(DebugLogLocation)
            if err != nil {
                log.Fatal("Error while parsing the log file", err)
            }
            for _, logEntry := range parsedLogs {
                if strings.Contains(logEntry.Event, "wave") {
                    discord.ChannelMessageSend(message.ChannelID, fmt.Sprintf("%v %s %s", logEntry.Timestamp, logEntry.Name, logEntry.Event))
                }
            }
    }
}

func Run() {
    discord, err := discordgo.New("Bot " + DiscordBotToken)
    if err != nil {
        log.Fatal(err)
    }
    discord.AddHandler(newMessage)
    discord.Open()
    defer discord.Close()
    log.Println(`Running...`)
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
}