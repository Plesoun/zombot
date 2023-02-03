package bot

import (
//    "bytes"
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
        case strings.Contains(message.Content, CommandPrefix + "help"):
            discord.ChannelMessageSend(message.ChannelID, "Available commands: \n !help \n !status \n !uptime \n !logouts")
        case strings.Contains(message.Content, CommandPrefix + "status"):
            discord.ChannelMessageSend(message.ChannelID, "Green")
        case strings.Contains(message.Content, CommandPrefix +"uptime"):
            out, err := exec.Command("bash", "-c", "sudo systemctl status docker | grep -Po '.*; \\K(.*)(?= ago)'").Output()
            if err != nil {
                log.Fatal("%s", err)
            }
            discord.ChannelMessageSend(message.ChannelID, string(out[:]))

        case strings.Contains(message.Content, CommandPrefix +"bot"):
            discord.ChannelMessageSend(message.ChannelID, "Zombot!")
        case strings.Contains(message.Content, CommandPrefix +"logouts"):
            for _, line := range parser.ReadLogFile("./sample_log.txt") {
                discord.ChannelMessageSend(message.ChannelID, line)
            }
    }
}

func Run() {
    // not implemented yet
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