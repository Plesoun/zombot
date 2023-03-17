package tests

import (
    "testing"
    "time"
    "zombot/parser"
)

type parsedLog struct {
    timestamp   time.Time
    name        string
    event       string
}

func TestParseLogLine(t *testing.T) {
    logLine := `[30-01-23 17:36:27.662] 76561197995472465 "Plesoun" fully connected (10156,6640,0).`

    expected := parsedLog{
        timestamp:  time.Date(2023, time.January, 30, 17, 36, 27, 0, time.UTC),
        name:       "Plesoun",
        event:      "fully connected",
    }

    result, err := parser.  ParseLogLine(logLine)
    if err != nil {
        t.Fatalf("Error while parsing line: %v", err)
    }

    if !result.timestamp.Equal(expected.timestamp) {
        t.Errorf("Expected timestamp: %v, got %v", expected.timestamp, result.timestamp)
    }
    if !result.name.Equal(expected.user) {
        t.Errorf("Expected name: %v, got %v", expected.name, result.name)
    }
    if !result.event.Equal(expected.event) {
        t.Errorf("Expected event: %v, got %v", expected.event, result.event)
    }
}
