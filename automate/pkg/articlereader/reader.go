package articlereader

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"robel-yemane.github.io/automate/pkg/types"
)

// Read reads paragraphs of text and returns the data read in a slice of strings.
func Read(r io.Reader) (*types.ArticleText, error) {

	var articleContent []string
	var title string

	reader := bufio.NewReader(r)
	for {
		bytesRead, err := reader.ReadBytes('\n')
		line := string(bytesRead)

		if _, after, found := strings.Cut(line, "Title:"); found {
			title = strings.TrimSpace(after)
			if err == io.EOF {
				break
			}
			continue
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("reading article: %w", err)
		}
		if line == "\n" {
			continue
		}

		articleContent = append(articleContent, strings.TrimSuffix(line, "\n"))
	}

	return &types.ArticleText{Title: title, Body: articleContent}, nil
}
