package server

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	ascii "server/functions"
)

type testCase struct {
	text        string
	banner      string
	expected    string
	expectError bool
}

var testCases = []testCase{
	{
		text:        "hello",
		banner:      "standard",
		expected:    Join("../functions/resources/test.txt"),
		expectError: false,
	},
}

func TestInput(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.text, func(t *testing.T) {
			// Capture stdout
			old := os.Stdout
			r, w, err := os.Pipe()
			if err != nil {
				fmt.Println("Error")
				return
			}
			os.Stdout = w

			// Execute the function
			output, err := ascii.Input(tc.text, tc.banner)
			if err != nil {
				fmt.Println("Error2")
				return
			}
			fmt.Println(output)

			// Close and restore stdout
			w.Close()
			os.Stdout = old
			var got bytes.Buffer
			_, err1 := got.ReadFrom(r)
			if err1 != nil {
				fmt.Println("Error3")
				return
			}
			// _, _ = io.Copy(&buf, r)

			// Check for errors if expected
			if tc.expectError {
				if err == nil {
					t.Fatalf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("didn't expect an error but got: %v", err)
				}

				// Compare output
				if got.String() != tc.expected {
					t.Errorf("expected %q, got %q", tc.expected, got)
				}
			}
		})
	}
}

func Join(filename string) string {
	text, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Test case file not found.")
	}

	return string(text)
}
