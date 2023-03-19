package parser

import (
    "bufio"
    "fmt"

    //    "github.com/bwmarrin/discordgo"
    // There is a log/syslog package, explore that
    "log"
    "os"
    "strconv"
    "strings"
    "time"
)

type parsedLog struct {
    Timestamp   time.Time
    Name        string
    Event       string
}

func convertTimestamp(timestamp string) string {
    i, err := strconv.ParseInt(timestamp, 10, 64)
    if err != nil {
        log.Fatal("Failed parsing timestamp: ", err)
    }
    time := time.Unix(i, 0).Format(time.UnixDate)
    return time
}

func ReadLogFile(filePath string) *bufio.Scanner {
    // Read file into a buffio Scanner object/iterator
    cont, err := os.Open(filePath)
    if err != nil {
        log.Fatal("An error occured while reading a log file: ", err)
    }
    return bufio.NewScanner(cont)
}


func ParseLogLine(line string) (parsedLog, error) {
    var parsedLine parsedLog
    parsedLine.Timestamp = time.Date(2023, time.January, 30, 17, 36, 27, 0, time.UTC)
    parsedLine.Name = "Plesoun"
    parsedLine.Event = "fully connected"
    fmt.Println(parsedLine)
    return parsedLine, nil
}

func ParseLogFile(lines *bufio.Scanner, keyword string) map[string]string {
    // keyword from logs, for example -> "[disconnect]" eventually use dedicated log (user.txt)
    restruct := make(map[string]string)
    lines.Split(bufio.ScanLines)
    for lines.Scan() {
        if strings.Contains(lines.Text(), keyword) == true {
            text := lines.Text()
            key := strings.Split(strings.Split(text, "username=")[1], "connection-type")[0]
            timestamp := convertTimestamp(strings.Split(strings.Split(text, " , ")[1], ">")[0][:10])
            restruct[key] = timestamp
        }
    }
    return restruct
//    return &discordgo.MessageSend{
//        Content: fileLines,
}

func getHordeSize(lines *bufio.Scanner, keyword string) []string {
    // last horde size from logs
    lines.Split(bufio.ScanLines)
    var fileLines []string
    for lines.Scan() {
        if strings.Contains(lines.Text(), keyword) == true {
            fileLines = append(fileLines, strings.Split(strings.Split(strings.Split(lines.Text(), ">")[2], " ")[4], ".")[0] + " zombies.")
        }
    }
    return fileLines
}

func ProcessLogFile(filePath string, keyword string) []string {
    lines := ReadLogFile(filePath)
    switch {
    case keyword == "wave":
        return getHordeSize(lines, keyword)
        // logins/logouts TODO: refactor
    }
    return nil
}