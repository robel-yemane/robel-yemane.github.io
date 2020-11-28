package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

func main() {
	const boilerphtml = `
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
		{{ range .Items -}}
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

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	t, err := template.New("blog").Parse(boilerphtml)
	check(err)

	// Parse a time value from a string in the standard Unix format.
	utime := time.Now() //time.Parse(time.Now .UnixDate, "Wed Feb 25 11:06:39 PST 2015")
	utimedate := utime.Format("2006-01-02")
	ftimedate := utime.Format("2006.01.02")

	data := struct {
		Title  string
		Header string
		Udate  string
		Fdate  string
		Items  []string
	}{
		Title:  "Robel Yemane",
		Header: "Lorem Ipsum",
		Udate:  utimedate,
		Fdate:  ftimedate,
		Items: []string{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque ex orci, consequat eget dignissim sed, tempus sed arcu. Integer semper odio massa, id convallis erat sollicitudin vitae. Nullam at massa non nunc accumsan suscipit. Nullam porta laoreet dui non sollicitudin. Nulla sapien enim, tincidunt sed fermentum in, varius vel quam. Suspendisse a molestie velit. Vestibulum rutrum finibus justo, in pretium elit fermentum eu. Integer aliquet sit amet mi non gravida. Etiam sit amet gravida sapien, nec aliquet magna. Integer tempor, est ac rhoncus vulputate, dui nunc interdum nulla, in scelerisque nisi augue ultrices mi. Aenean congue nisl blandit nulla volutpat vulputate. Donec mollis justo quis tincidunt porta.",
			"Duis laoreet felis sit amet varius hendrerit. Nunc tincidunt, ex a convallis tristique, metus neque maximus tellus, in elementum elit ante et neque. Sed in ultricies diam, vel molestie nisi. Pellentesque quam justo, porta eu nisi vitae, facilisis sodales lectus. Sed condimentum, odio interdum mollis consectetur, neque mi finibus metus, at eleifend felis ante vitae augue. Fusce nec ante eget elit dignissim lobortis non ac nisl. Quisque vitae lectus vel erat aliquam fermentum ac sit amet metus.",
		},
	}
	// print out html file
	log.Println("Writing html file.")
	err = writeToFile(t, data)
	check(err)
	log.Println("Wrote html file.")

}

func writeToFile(t *template.Template, data interface{}) error {
	file, err := os.Create("article.html")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}

	err = t.Execute(file, data)
	return err
}

//TODO:
/*
	- read items from stdin/user
	- loops through contact section
	- write tests!!!!
	- github action wire!
*/
