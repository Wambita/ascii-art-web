package server

import (
	"fmt"
	"strings"
	"testing"

	ascii "server/ascii"
)

func TestAsciiCombine(t *testing.T) {
	input := "hello"
	expected := Join("../ascii/resources/test5.txt")

	result, err := ascii.AsciiCombine(input, ascii.AsciiArtMap(Join("../ascii/resources/standard.txt")))
	if err != nil {
		fmt.Println("failed")
	}
	if strings.Join(result,"") != expected {
		t.Errorf("%s\n%s\n", expected, result)
	}
}
