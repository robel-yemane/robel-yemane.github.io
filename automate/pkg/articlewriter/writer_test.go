package articlewriter

import (
	"bytes"
	"log"
	"testing"
	"time"

	"golang.org/x/net/html"
)

var body = []string{"Lorem ipsum dolor sit amet,", "consectetur adipiscing elit."}

const boilerPHTML = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
		<title>{{.Title}}</title>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" type="text/css" href="../styles/style.css">
    </head>
	<body>
		<header>
			<h1><a target="_self" href="../index.html">{{.Title}}</a></h1>
  		</header>
		<section id="content">
		<h2>{{.Header}}</h2>
		<time datetime="{{.Udate}}">{{.Fdate}}</time>
		{{ range .Body -}}
		<p>
		 {{ . }}
		</p>
		{{ end -}}
		</section>
    	<section id="contact">
      		<ul>
        		<li><a target="_self" href="https://www.linkedin.com/in/ryemane/">linkedin</a></li>
      		</ul>
    	</section>
    </body>
</html>`

var utimedate = time.Now().Format("2006-01-02")
var ftimedate = time.Now().Format("2006.01.02")

var htmlLayout = struct {
	Title    string
	Header   string
	Udate    string
	Fdate    string
	Body     []string
	Email    string
	Twitter  string
	Linkedin string
}{
	Title:  "Robel Yemane",
	Header: "Lorem Ipsum",
	Udate:  utimedate,
	Fdate:  ftimedate,
	Body:   body,
}

// https://pkg.go.dev/golang.org/x/net/html
func TestWrite(t *testing.T) {

	var bufWriter bytes.Buffer //=> has Read and Write method
	err := Write(boilerPHTML, body, &bufWriter)

	if err != nil {
		t.Fatal("Could not Write html output", err)
	}

	doc, err := html.Parse(&bufWriter)
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	var counter int
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			if n.Data == htmlLayout.Title || n.Data == htmlLayout.Header {
				counter++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if counter != 3 {
		t.Errorf("Parsed html output does not contain expected data")
	}
}
