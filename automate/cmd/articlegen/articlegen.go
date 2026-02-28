package main

import (
	"log"
	"os"

	flag "github.com/spf13/pflag"
	"robel-yemane.github.io/automate/pkg/articlereader"
	"robel-yemane.github.io/automate/pkg/articlewriter"
	"robel-yemane.github.io/automate/pkg/types"
)

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
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("getting working directory: %v", err)
	}

	var srcArticlePath string
	var outHTMLPath string
	var email string
	var twitter string
	var linkedin string

	flag.StringVarP(&srcArticlePath, "article", "a", "", "[required] Full path to the source article")
	flag.StringVarP(&outHTMLPath, "htmlOut", "o", cwd+"/article.html", "Full path to the out html file.")
	flag.StringVar(&email, "email", "", "Contact email address")
	flag.StringVar(&twitter, "twitter", "", "Twitter profile URL")
	flag.StringVar(&linkedin, "linkedin", "", "LinkedIn profile URL")
	flag.Parse()

	if srcArticlePath == "" {
		log.Println("Source `article` flag must be provided.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	tmpl, err := articlewriter.ParseTemplate(boilerPHtml)
	if err != nil {
		log.Fatalf("parsing template: %v", err)
	}

	srcFile, err := os.Open(srcArticlePath)
	if err != nil {
		log.Fatalf("opening article %q: %v", srcArticlePath, err)
	}
	defer srcFile.Close()

	articleContent, err := articlereader.Read(srcFile)
	if err != nil {
		log.Fatalf("reading article: %v", err)
	}

	log.Println("Writing html file.")

	outFile, err := os.Create(outHTMLPath)
	if err != nil {
		log.Fatalf("creating output file %q: %v", outHTMLPath, err)
	}
	defer outFile.Close()

	contact := types.Contact{Email: email, Twitter: twitter, Linkedin: linkedin}
	if err := articlewriter.Write(tmpl, *articleContent, contact, outFile); err != nil {
		log.Fatalf("writing html: %v", err)
	}

	log.Printf("Wrote html file: [%s]", outHTMLPath)
}
