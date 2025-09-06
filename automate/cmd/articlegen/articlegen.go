package main

import (
	"log"
	"os"
	"path/filepath"

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
	tmplPath := filepath.Join(filepath.Dir(filepath.Dir(os.Args[0])), "templates", "article.html")
	err = articlewriter.WriteFromTemplate(tmplPath, *articleContent, file)
	check(err)

	log.Printf("Wrote html file: [%s]", outHTMLPath)

}
