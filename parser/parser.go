package parser

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
    "time"
)

func convertTimestamp(timestamp string) string {
    i, err := strconv.ParseInt(timestamp, 10, 64)
    if err != nil {
        log.Fatal("Failed parsing timestamp: ", err)
    }
    time := time.Unix(i, 0)
    return time.Format("2023-01-02 15:04:05")
}

func readLogFile(filePath string) *bufio.Scanner {
    // Read file into a buffio Scanner object/iterator
    cont, err := os.Open(filePath)
    if err != nil {
        log.Fatal("An error occured while reading a log file: ", err)
    }
    return bufio.NewScanner(cont)
}

func parseLogFile(lines *bufio.Scanner, keyword string) []string {
    // keyword from logs, for example -> "[disconnect]"
    lines.Split(bufio.ScanLines)
    var fileLines []string
    for lines.Scan() {
        if strings.Contains(lines.Text(), keyword) == true {
            fileLines = append(fileLines, strings.Split(lines.Text(), "] ")[0] + "] " + strings.Split(lines.Text(), "username=")[1])
        }
    }
    return fileLines
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
    fmt.Print(fileLines)
    return fileLines
}

func ProcessLogFile(filePath string, keyword string) []string {
    lines := readLogFile(filePath)
    switch {
    case keyword == "wave":
        return getHordeSize(lines, keyword)
    // logins/logouts TODO: refactor
    case strings.Contains(keyword, "["):
        return parseLogFile(lines, keyword)
    }
    return nil
}