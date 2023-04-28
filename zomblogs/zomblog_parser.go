package parser

import (
    "bufio"
    "errors"
    "fmt"
    //    "github.com/bwmarrin/discordgo"
    // There is a log/syslog package, explore that
    "log"
    "os"
    "strings"
    "time"
)

type ParsedLog struct {
    Timestamp   time.Time
    Name        string
    Event       string
}

func ReadLogFile(filePath string) *bufio.Scanner {
    // Read file into a buffio Scanner object/iterator
    cont, err := os.Open(filePath)
    if err != nil {
        log.Fatal("An error occured while reading a log file: ", err)
    }
    return bufio.NewScanner(cont)
}


func ParseLogLine(line string) (ParsedLog, error) {
    var parsedLine ParsedLog
    // Find timestamp
    startIndex := strings.Index(line, "[")
    endIndex := strings.Index(line, "]")
    if startIndex == -1 || endIndex == -1 {
        return parsedLine, errors.New("invalid log format (timestamp)")
    }
    timestampStr := line[startIndex+1 : endIndex]
    fmt.Println(timestampStr)
    timestamp, err := time.Parse("02-01-06 15:04:05", timestampStr[:len(timestampStr)-4])
    if err != nil {
        return parsedLine, errors.New("invalid timestamp")
    }
    parsedLine.Timestamp = timestamp

    // Find connection details
    startIndex = strings.Index(line, "\"")
    endIndex = strings.LastIndex(line, "\"")
    if startIndex == -1 || endIndex == -1 || startIndex == endIndex {
        return parsedLine, errors.New("invalid log format (name)")
    }
    parsedLine.Name = line[startIndex+1 : endIndex]
    if strings.Contains(line, "fully connected") {
        parsedLine.Event = "fully connected"
    } else {
       return parsedLine, errors.New("event not found")
    }
    return parsedLine, nil
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

func ParseLogFile(filePath string) []string {
    lines := ReadLogFile(filePath)
    lines.Split(bufio.ScanLines)
    for lines.Scan() {
        ParseLogLine(lines.Text())
    }
    return nil
}