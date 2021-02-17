package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
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

func createPDFfromTextFile(textFilePath string) {
	// fmt.Println(wkhtmltopdf.GetPath())
	// wkhtmltopdf.SetPath("./")
	// fmt.Println(wkhtmltopdf.GetPath())
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(textFilePath)
	htmlfile, err := ioutil.ReadFile(textFilePath)
	if err != nil {
		log.Fatal(err)
	}
	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(htmlfile)))
	pdfg.Dpi.Set(600)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}
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

	var textFilePathHTML string
	var textFilePathPDF string
	var directory string

	flag.StringVar(&textFilePathHTML, "file", "", "Render text file to HTML.")
	flag.StringVar(&textFilePathPDF, "pdf", "", "Render HTML file to PDF.")
	flag.StringVar(&directory, "dir", "", "Find all .txt files in the given directory.")
	flag.Parse()

	if textFilePathHTML != "" {
		createPageFromTextFile(textFilePathHTML)
		// test()
	}

	if textFilePathPDF != "" {
		// fmt.Println(textFilePathPDF)
		createPDFfromTextFile(textFilePathPDF)
	}

	if directory != "" {
		textFiles := findAllFilesInDirectory(directory)
		for _, fileName := range textFiles {
			createPageFromTextFile(fileName)
		}
	}

	// flag.Visit(func(f *flag.Flag) {
	// 	if f.Name == "pdf" {
	// 		fmt.Print("I PASS")
	// 	}
	// })

}
