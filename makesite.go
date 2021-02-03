package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

// Update the `save` function to use the input filename to generate a new HTML file.
func save(textFilePtr *string) {
	textFile, err := ioutil.ReadFile(*textFilePtr)
	textFileName := strings.Split(*textFilePtr, ".")[0]
	textToHTML, err := os.Create(fmt.Sprintf("%s.html", textFileName))
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(textToHTML, textFile)
	if err != nil {
		panic(err)
	}
}

func main() {
	// Read in the contents of the provided `first-post.txt` file.
	// firstPostTxt, err := ioutil.ReadFile("first-post.txt")

	// Render the contents of `first-post.txt` using Go Templates and print it to stdout.
	// t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	// err = t.Execute(os.Stdout, string(firstPostTxt))
	// if err != nil {
	// 	panic(err)
	// }

	// Write the HTML template to the filesystem to a file. Name it `first-post.html`.
	// myFile, err := os.Create("first-post.html")
	// err = t.Execute(myFile, string(firstPostTxt))
	// if err != nil {
	// 	panic(err)
	// }

	// Add a new flag to your command named `file`. This flag represents the name of
	// any `.txt` file in the same directory as your program.
	textFilePtr := flag.String("file", "", "Text file to render to HTML.")
	flag.Parse()
	save(textFilePtr)

}
