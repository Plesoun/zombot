package parser

import (
    "bufio"
    "errors"
    "fmt"
    // There is a log/syslog package, explore that
    "os"
    "strings"
    "time"
)

type ParsedLog struct {
    Timestamp   time.Time
    Name        string
    Event       string
}

func ReadLogFile(filePath string) (*bufio.Scanner, *os.File, error) {
    // Read file into a buffio Scanner object/iterator.
    //!! This means caller is responsible for closing the file
    cont, err := os.Open(filePath)
    if err != nil {
        return nil, nil, fmt.Errorf("an error occured while reading the log file: %w", err)
    }
    return bufio.NewScanner(cont), cont, nil
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
    timestamp, err := time.Parse("02-01-06 15:04:05", timestampStr[:17])
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
    // Extracting the player based event here
    eventStartIndex := endIndex + 2
    eventEndIndex := strings.LastIndex(line, "(")
    if eventStartIndex < len(line) {
        parsedLine.Event = line[eventStartIndex:eventEndIndex -1]
    } else {
       return parsedLine, errors.New("event not found")
    }
    return parsedLine, nil
}

func ParseLogFile(filePath string) ([]ParsedLog, error) {
    lines, cont, err := ReadLogFile(filePath)
    if err != nil {
        return nil, err
    }
    defer cont.Close()

    var parsedLogs []ParsedLog
    for lines.Scan() {
        parsedLogline, err := ParseLogLine(lines.Text())
        fmt.Println(parsedLogline)
        if err != nil {
            //handle error here (maybe log, maybe raise)
            continue
        }
        parsedLogs = append(parsedLogs, parsedLogline)
    }
    if err := lines.Err(); err != nil {
        return nil, err
    }
    return parsedLogs, nil
}