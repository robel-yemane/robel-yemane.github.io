package articlereader

import (
	"bufio"

	"io"
	"log"
	"strings"

	"robel-yemane.github.io/automate/pkg/types"
)

// Read reads paragraphs of text and returns the data read
// in a slice of strings
func Read(r io.Reader) *types.ArticleText {

	articleContent := []string{}
	var title string

	reader := bufio.NewReader(r)
	for {
		bytesRead, err := reader.ReadBytes('\n')

		if strings.Contains(string(bytesRead), "Title:") {
			titleRunes := []rune(string(bytesRead))

			title = string(titleRunes[6:]) // fetch the title and convert it back to string
			continue                       //TODO: is there a better alternative to skipping processing the first line after the first run?
		}

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

	a := types.ArticleText{Title: title, Body: articleContent}

	return &a
}
