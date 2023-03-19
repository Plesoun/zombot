package tests

import (
    "testing"
    "time"
    "zombot/zomblogs"
)

type parsedLog struct {
    Timestamp   time.Time
    Name        string
    Event       string
}

func TestParseLogLine(t *testing.T) {
    logLine := `[30-01-23 17:36:27.662] 76561197995472465 "Plesoun" fully connected (10156,6640,0).`

    expected := parsedLog{
        Timestamp:  time.Date(2023, time.January, 30, 17, 36, 27, 0, time.UTC),
        Name:       "Plesoun",
        Event:      "fully connected",
    }

    result, err := parser.ParseLogLine(logLine)
    if err != nil {
        t.Fatalf("Error while parsing line: %v", err)
    }

    if !result.Timestamp.Equal(expected.Timestamp) {
        t.Errorf("Expected timestamp: %v, got %v", expected.Timestamp, result.Timestamp)
    }
    if result.Name != expected.Name {
        t.Errorf("Expected name: %v, got %v", expected.Name, result.Name)
    }
    if result.Event != expected.Event {
        t.Errorf("Expected event: %v, got %v", expected.Event, result.Event)
    }
}
