package parser

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func ReadLogFile(filePath string) {
    cont, err := os.Open(filePath)
    if err != nil {
        log.Fatal("An error occured while reading a log file: ", err)
    }
    lines := bufio.NewScanner(cont)
    lines.Split(bufio.ScanLines)
    var fileLines []string
    for lines.Scan() {
        fileLines = append(fileLines, lines.Text())
    }
    cont.Close()

    for _, line := range fileLines {
        fmt.Println(line)
    }
}
