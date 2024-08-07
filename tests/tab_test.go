package server

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	ascii "server/ascii"
)

type testTab struct {
	text          string
	banner        string
	expected      string
	expectedError bool
}

var testTabs = []testTab{
	{
		text:          "hello\tthere ",
		banner:        "standard",
		expected:      Join("../ascii/resources/test7.txt"),
		expectedError: false,
	},
}

func TestTab(t *testing.T) {
	for _, tc := range testTabs {
		t.Run(tc.text, func(t *testing.T) {
			old := os.Stdout
			r, w, err := os.Pipe()
			if err != nil {
				fmt.Println("Error4")
				return
			}

			os.Stdout = w

			output, err := ascii.Input(tc.text, tc.banner)
			if err != nil {
				fmt.Println("Error5")
				return
			}
			fmt.Println(output)

			w.Close()
			os.Stdout = old
			var got bytes.Buffer
			_, err1 := got.ReadFrom(r)
			if err1 != nil {
				fmt.Println("Error6")
				return
			}

			if got.String() != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, got)
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