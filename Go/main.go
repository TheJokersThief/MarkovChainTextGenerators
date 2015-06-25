package main

import (
	// "fmt"
	"io"
	"os"
	"path/filepath"
	// "io/ioutil"
	"bytes"
	"strings"
)

// Performs the main test execution of the program
func main() {
	
	markov := Markov{ }
    markov.openFile( )
    markov.splitWords( )
}

// A 2nd Order Markov Chain Text Generator
type Markov struct{
	store map[string][]string
	words []string
	fileContents string
}

// Opens the file for reading and converts it from bytes
// to a string
func (markov *Markov) openFile() {
	// Get the absolute path to open our file
	absPath, _ := filepath.Abs("../test_text.txt")
    buffer := bytes.NewBuffer(nil)
	file, err := os.Open(absPath)
	if err != nil {
	    panic(err)
	}
	// Copy our file into a bytes buffer
	io.Copy(buffer, file)           
	file.Close()

	// Convert those bytes into a massive string
	fileContents := string(buffer.Bytes())
	markov.fileContents = fileContents

}

// Splits all of the words up into a string slice
func (markov *Markov) splitWords( ){
	markov.words = strings.Split(markov.fileContents, " ")
}