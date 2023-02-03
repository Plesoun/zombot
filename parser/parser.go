package parser

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

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
    fmt.Println(fileLines)
    return fileLines
}

func ProcessLogFile(filePath string, keyword string) []string {
    lines := readLogFile(filePath)
    return parseLogFile(lines, keyword)
}