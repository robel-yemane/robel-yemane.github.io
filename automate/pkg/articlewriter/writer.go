package articlewriter

import (
	"html/template"
	"io"
	"time"

	"robel-yemane.github.io/automate/pkg/types"
)

// ParseTemplate parses the HTML skeleton once and returns a ready-to-use template.
func ParseTemplate(htmlSkeleton string) (*template.Template, error) {
	return template.New("blog").Parse(htmlSkeleton)
}

// Write executes the pre-parsed template with article data and contact info.
func Write(tmpl *template.Template, data types.ArticleText, contact types.Contact, w io.Writer) error {
	utime := time.Now()

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
		Udate:    utime.Format("2006-01-02"),
		Fdate:    utime.Format("2006.01.02"),
		Body:     data.Body,
		Email:    contact.Email,
		Twitter:  contact.Twitter,
		Linkedin: contact.Linkedin,
	}

	return tmpl.Execute(w, htmlLayout)
}
