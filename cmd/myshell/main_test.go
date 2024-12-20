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
		{"cd /home", command{name: "cd", args: []string{"/home"}}},
		{"echo 'hello world'", command{name: "echo", args: []string{"hello world"}}},
		{"type echo", command{name: "type", args: []string{"echo"}}},
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
