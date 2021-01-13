package articlereader

import (
	"bufio"
	"io"
	"log"
	"strings"
)

// Read reads paragraphs of text and returns the data read
// in a slice of strings
func Read(r io.Reader) []string {

	articleContent := []string{}
	reader := bufio.NewReader(r)
	for {
		bytesRead, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if string(bytesRead) == "\n" {
			continue
		}
		// trim the '\n' returned from ReadBytes()
		articleContent = append(articleContent, strings.TrimSuffix(string(bytesRead), "\n"))
	}

	return articleContent
}
