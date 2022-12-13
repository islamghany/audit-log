package logger

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

func TestLogger(t *testing.T) {

	tt := []struct {
		message     string
		properties  map[string]string
		trace       string
		description string
		level       Level
	}{
		{"Hello there", nil, "", "Test With properties", LevelInfo},
		{"Printing something useful", map[string]string{"name": "logger"}, "", "Test With properties", LevelInfo},
		{"Hello there", nil, "", "Test With properties", LevelError},
		{"Printing something useful", map[string]string{"name": "logger"}, "", "Test With properties", LevelError},
	}

	var output bytes.Buffer
	l := New(&output, LevelInfo)

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			switch tc.level {
			case LevelInfo:
				l.PrintInfo(tc.message, tc.properties)
			case LevelError:
				l.PrintError(errors.New(tc.message), tc.properties)

			default:
				t.Fatal("not expcted level value")
			}
			var aux ILogger

			err := json.Unmarshal(output.Bytes(), &aux)

			if err != nil {
				t.Fatal(err)
			}

			if tc.message != aux.Message {
				t.Errorf("got %s but expected %s", aux.Message, tc.message)
			}
			if tc.properties != nil {
				isSame := reflect.DeepEqual(tc.properties, aux.Properties)
				if !isSame {
					t.Errorf("got %+v but expected %+v", aux.Properties, tc.properties)
				}

			} else {
				if aux.Properties != nil {
					t.Errorf("expected properites to be nil but got %v", aux.Properties)
				}

			}
			if tc.level == LevelError && aux.Trace == "" {
				t.Error("expected trace have a value")
			}
			if tc.level == LevelInfo && aux.Trace != "" {
				t.Errorf("expected trace to be empty but got %s", aux.Trace)
			}
			output.Reset()
		})
	}

}
