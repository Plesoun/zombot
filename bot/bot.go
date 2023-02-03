package bot

import (
    "github.com/bwmarrin/discordgo"
    "log"
    "os"
    "os/exec"
    "os/signal"
    "strings"
    "zombot/parser"
)

var (
    DiscordBotToken string
    CommandPrefix string
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
            discord.ChannelMessageSend(message.ChannelID, "Available commands: \n !help \n !status \n !logouts \n !logins \n !system")
        // system returns basic system stats from "top"
        case strings.Contains(message.Content, CommandPrefix +"system"):
            out, err := exec.Command("bash", "-c", "top -b -n 1 | egrep 'top -|Tasks:|%Cpu|MiB'").Output()
            if err != nil {
                log.Fatal("Encountered error while executing system command: ", err)
            }
            discord.ChannelMessageSend(message.ChannelID, string(out[:]))
        // status returns the systemd status section relevant to the query
        case strings.Contains(message.Content, CommandPrefix +"status"):
            out, err := exec.Command("bash", "-c", "sudo systemctl status docker | grep Active").Output()
            if err != nil {
                log.Fatal("Encountered error while executing system command: ", err)
            }
            discord.ChannelMessageSend(message.ChannelID, string(out[:]))
        // see all logouts contained in the current log file TODO: limit to ~ last 10
        case strings.Contains(message.Content, CommandPrefix +"logouts"):
            for _, line := range parser.ProcessLogFile("./sample_log.txt", "[disconnect]") {
                discord.ChannelMessageSend(message.ChannelID, line)
            }
        // see all logins contained in the current log file TODO: limit to ~ last 10
        case strings.Contains(message.Content, CommandPrefix +"logins"):
            for _, line := range parser.ProcessLogFile("./sample_log.txt", "[fully-connected]") {
                discord.ChannelMessageSend(message.ChannelID, line)
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
    log.Println("Running...")
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
}