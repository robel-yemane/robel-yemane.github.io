package articlereader

import (
	"strings"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestRead(t *testing.T) {
	article := "Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit.Pellentesque ex orci, consequat eget\n dignissim sed,tempus sed arcu. Integer semper odio massa\n"
	stringReader := strings.NewReader(article)
	t.Log("\tGiven the need to test reading paragraphs.")

	content := Read(stringReader)
	if content == nil {
		t.Fatal("\t\tShould be able to read the contents of the file.", ballotX, content)
	}
	t.Log("\t\tShould be able to read the contents of the file.", checkMark)
	paragraphs := len(content)
	lineNum := 3
	if paragraphs == lineNum {
		t.Logf("\t\tFile should have %d lines. %v", lineNum, checkMark)
	} else {
		t.Errorf("\t\tFile should have %d lines. %v", lineNum, ballotX)
	}
	articleParagraphs := strings.Split(article, "\n")

	if strings.Trim(content[1], "\n") != strings.Trim(articleParagraphs[1], "\n") {
		t.Errorf("\t\t First paragraph read should be\n%v\n - Found: \n%v\n", content[1], articleParagraphs[1])
	}

}
