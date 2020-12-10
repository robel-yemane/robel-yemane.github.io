package main

import (
	"log"
	"os"

	"robel-yemane.github.io/automate/articlereader"
	"robel-yemane.github.io/automate/articlewriter"
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
        		<li><a target="_self" href="https://twitter.com/robelyemane_">twitter</a></li>
        		<li><a target="_self" href="mailto:iamsnowball3@gmail.com">e-mail</a></li>
        		<li><a target="_self" href="https://www.linkedin.com/in/ryemane/">linkedin</a></li>
      		</ul>
    	</section>
    </body>
</html>`

func main() {

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	// TODO: read these from user/stdin
	plainArticle := "article.txt"
	htmlContent := "article.html"

	file, err := os.Open(plainArticle)
	check(err)

	//TODO: should err be returned from Read and
	// and handled here?
	//read file contents
	articleContent := articlereader.Read(file)

	//write file contents into html file
	log.Println("Writing html file.")

	file, err = os.Create(htmlContent)
	err = articlewriter.Write(boilerPHtml, articleContent, file)
	check(err)
	log.Println("Wrote html file.")

}

//TODO:
/*
	- read items from file
	- loops through contact section
	- write tests!!!!
	- github action wire!
*/
