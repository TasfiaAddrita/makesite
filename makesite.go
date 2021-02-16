package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type HTMLPage struct {
	Title   string
	Content string
}

// Create a function to use the input filename to generate a new HTML file.
func createPageFromTextFile(textFilePath string) {
	textFile, err := ioutil.ReadFile(textFilePath)
	textFileName := strings.Split(textFilePath, ".")[0]
	htmlFile, err := os.Create(fmt.Sprintf("%s.html", textFileName))

	textToHTML := HTMLPage{textFileName, string(textFile)}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(htmlFile, textToHTML)
	if err != nil {
		panic(err)
	}
}

func test() {
	fmt.Print("I am test")
}

func findAllFilesInDirectory(directory string) []string {
	var textFiles []string

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := strings.Split(file.Name(), ".")
		if len(fileName) > 1 && fileName[1] == "txt" {
			// fmt.Println(fileName)
			textFiles = append(textFiles, file.Name())
		}
	}
	return textFiles
}

func main() {

	var textFilePath string
	var directory string

	flag.StringVar(&textFilePath, "file", "", "Render text file to HTML.")
	flag.StringVar(&directory, "dir", "", "Find all .txt files in the given directory.")
	flag.Parse()

	if textFilePath != "" {
		createPageFromTextFile(textFilePath)
		// test()
	}

	if directory != "" {
		textFiles := findAllFilesInDirectory(directory)
		for _, fileName := range textFiles {
			createPageFromTextFile(fileName)
		}
	}

}
