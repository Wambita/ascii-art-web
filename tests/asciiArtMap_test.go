package server

import (
	"testing"

	ascii "server/ascii"
)

func TestAsciiArt(t *testing.T) {
	input := "hi"
	expected := Join("../ascii/resources/hi.txt")

	result, _ := ascii.Art(input, ascii.AsciiArtMap(Join("../ascii/resources/standard.txt")))
	// fmt.Println(result)
	if result != expected {
		t.Errorf("Got %q expected %q", result, expected)
	}
}
