package articlewriter

import (
	"html/template"
	"io"
	"time"
)

// Write writes html formatted data into a file
func Write(htmlSkeleton string, data []string, w io.Writer) error {

	// Parse a time value from a string in the standard Unix format.
	utime := time.Now()
	utimedate := utime.Format("2006-01-02")
	ftimedate := utime.Format("2006.01.02")
	email := "ryhgb03@gmail.com"
	twitter := "https://twitter.com/robelyemane_"
	linkedin := "https://www.linkedin.com/in/ryemane/"

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
		Header:   "Lorem Ipsum",
		Udate:    utimedate,
		Fdate:    ftimedate,
		Body:     data,
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
