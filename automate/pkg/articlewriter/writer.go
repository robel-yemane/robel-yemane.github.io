package articlewriter

import (
	"html/template"
	"io"
	"time"

	"robel-yemane.github.io/automate/pkg/types"
)

var email = "ryhgb03@gmail.com"
var twitter = "https://twitter.com/robelyemane_"
var linkedin = "https://www.linkedin.com/in/ryemane/"

// Write writes html formatted data into a file
func Write(htmlSkeleton string, data types.ArticleText, w io.Writer) error {

	// Parse a time value from a string in the standard Unix format.
	utime := time.Now()
	utimedate := utime.Format("2006-01-02")
	ftimedate := utime.Format("2006.01.02")

	htmlLayout := struct {
		Title    string
		Header   string
		Udate    string
		Fdate    string
		Body     []string
		Email    string
		Twitter  string
		Linkedin string
	}{
		Title:    "Robel Yemane",
		Header:   data.Title,
		Udate:    utimedate,
		Fdate:    ftimedate,
		Body:     data.Body,
		Email:    email,
		Twitter:  twitter,
		Linkedin: linkedin,
	}
	t, err := template.New("blog").Parse(htmlSkeleton)
	if err != nil {
		return err
	}
	return t.Execute(w, htmlLayout)
}
