package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected command
	}{
		{"echo hello", command{name: "echo", args: []string{"hello"}}},
		{"exit", command{name: "exit", args: []string{}}},
		{"echo 'hello world'", command{name: "echo", args: []string{"hello world"}}},
		{"echo 'hello     script'", command{name: "echo", args: []string{"hello     script"}}},
		{"echo script     world", command{name: "echo", args: []string{"script", "world"}}},
		{"cat '/tmp/file name' '/tmp/file name with spaces'", command{name: "cat", args: []string{"/tmp/file name", "/tmp/file name with spaces"}}},
	}

	for _, test := range tests {
		result := parse(test.input)
		t.Run(test.input, func(t *testing.T) {
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}
