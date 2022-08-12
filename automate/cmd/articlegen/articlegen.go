package main

import (
	"log"
	"os"

	flag "github.com/spf13/pflag"
	"robel-yemane.github.io/automate/pkg/articlereader"
	"robel-yemane.github.io/automate/pkg/articlewriter"
)

var srcArticlePath string
var outHTMLPath string

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	flag.StringVarP(&srcArticlePath, "article", "a", "", "[required] Full path to the source article")
	flag.StringVarP(&outHTMLPath, "htmlOut", "o", path+"/article.html", "Full path to the out html file.")
}

const boilerPHtml = `
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
        		<li><a target="_self" href="{{.Twitter}}">twitter</a></li>
        		<li><a target="_self" href="mailto:{{.Email}}">e-mail</a></li>
        		<li><a target="_self" href="{{.Linkedin}}">linkedin</a></li>
      		</ul>
    	</section>
    </body>
</html>`

func main() {

	flag.Parse()

	if srcArticlePath == "" {
		log.Println("Source `article` flag must be provided.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.Open(srcArticlePath)
	check(err)

	//read file contents
	articleContent := articlereader.Read(file)

	log.Println("Writing html file.")

	// create file
	file, err = os.Create(outHTMLPath)
	check(err)
	//write file contents into html file
	err = articlewriter.Write(boilerPHtml, articleContent, file)
	check(err)

	log.Printf("Wrote html file: [%s]", outHTMLPath)

}
