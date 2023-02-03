package parser

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func ReadLogFile(filePath string) []string {
    cont, err := os.Open(filePath)
    if err != nil {
        log.Fatal("An error occured while reading a log file: ", err)
    }
    lines := bufio.NewScanner(cont)
    lines.Split(bufio.ScanLines)
    var fileLines []string
    for lines.Scan() {
        if strings.Contains(lines.Text(), "[disconnect]") == true {
            fileLines = append(fileLines, strings.Split(lines.Text(), "] ")[0] + "] " + strings.Split(lines.Text(), "username=")[1])
        }
    }
    cont.Close()
    fmt.Println(fileLines)
    return fileLines
//
//    for _, line := range fileLines {
//        fmt.Println(line)
//    }
}
