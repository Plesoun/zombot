package tests

import (
    "errors"
    "testing"
    "time"
    "zombot/zomblogs"
)

func TestParseLogLine(t *testing.T) {
    testCases := []struct {
        logLine         string
        expectedLog     parser.ParsedLog
        expectedError   error
    }{
        {
            logLine:        `[30-01-23 17:36:27.662] 76561197995472465 "Plesoun" fully connected (10156,6640,0).`,
            expectedLog:    parser.ParsedLog{
                Timestamp:  time.Date(2023, time.January, 30, 17, 36, 27, 0, time.UTC),
                Name:       "Plesoun",
                Event:      "fully connected",
            },
            expectedError:  nil,
        },
        {
            logLine:        `[30-01-23 17:36:27.662] 76561197995472465 "Plesoun" connected (10156,6640,0).`,
            expectedError:  errors.New("event not found"),
        },
        {
            logLine:        `30-01-23 17:36:27.662] 76561197995472465 "Plesoun" fully connected (10156,6640,0).`,
            expectedError:  errors.New("invalid log format (timestamp)"),
        },
    }

    for _, testCase := range testCases {
        result, err := parser.ParseLogLine(testCase.logLine)
        if err != nil {
            if testCase.expectedError == nil || err.Error() != testCase.expectedError.Error() {
                t.Fatalf("Unexpected error: %v", err)
            }
            // Error occurred and was expected, continue with the next test case
            continue
        }

        if testCase.expectedError != nil {
            t.Fatalf("Expected error %v, got no error", testCase.expectedError)
        }

        if !result.Timestamp.Equal(testCase.expectedLog.Timestamp) {
            t.Errorf("Expected timestamp %v, got %v", testCase.expectedLog.Timestamp, result.Timestamp)
        }

        if result.Name != testCase.expectedLog.Name {
            t.Errorf("Expected name %s, got %s", testCase.expectedLog.Name, result.Name)
        }

        if result.Event != testCase.expectedLog.Event {
            t.Errorf("Expected event %s, got %s", testCase.expectedLog.Event, result.Event)
        }
    }
}
