package main

import (
	// "fmt"
	"html/template"
	"io/ioutil"
	"os"
)

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

func main() {
	// READING A FILE
	// fileContents, err := ioutil.ReadFile("first-post.txt")
	// if err != nil {
	// 	// A common use of `panic` is to abort if a function returns an error
	// 	// value that we don’t know how to (or want to) handle. This example
	// 	// panics if we get an unexpected error when creating a new file.
	// 	panic(err)
	// }
	// fmt.Print(string(fileContents))

	// WRITING A FILE
	// bytesToWrite := []byte("hello\ngo\n")
	// err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	// if err != nil {
	// 	panic(err)
	// }

	// Read in the contents of the provided `first-post.txt` file.
	fileContents, err := ioutil.ReadFile("first-post.txt")

	// Render the contents of `first-post.txt` using Go Templates and print it to stdout.
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(os.Stdout, string(fileContents))
	if err != nil {
		panic(err)
	}

	// Write the HTML template to the filesystem to a file. Name it `first-post.html`.
	myFile, err := os.Create("first-post.html")
	err = t.Execute(myFile, string(fileContents))
	if err != nil {
		panic(err)
	}
}
