package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"bytes"
	"strings"
	"math/rand"
	"time"
)

// Performs the main test execution of the program
func main() {
	
	markov := Markov{ }
    markov.openFile( )
    markov.splitWords( )
    markov.storeProbability( )
    text := markov.generateMarkovText( 250 )
    fmt.Println( text )
}

// A 2nd Order Markov Chain Text Generator
type Markov struct{
	store map[string][]string
	words []string
	fileContents string
	keys []string
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

// Creates a simple 2nd Order probability table for the words
func (markov *Markov) storeProbability( ){
	markov.store = make( map[string][]string )

	wordSize := len(markov.words)

	if wordSize < 3{
		// If there's less than three words, don't bother
		return
	} else {
		for i := 0; i < wordSize - 2; i++ {
			// Uses first and second word as a key, simulating probablity
			// by appending all succeeding words to an array
			key := markov.words[ i ] + ">>" + markov.words[ i + 1 ]
			markov.keys = append( markov.keys, key )
			markov.store[ key ] = append( markov.store[ key ], markov.words[ i + 2 ] )
		}
	}
}

// Generate the markov chain text based on the calculated store
func (markov Markov) generateMarkovText( wordLimit int ) string{
	output := ""
	keysLength := len( markov.keys )

	if !( len(markov.store) == 0 ){
		// Make sure the store isn't empty before beginning

		rand.Seed( time.Now().UnixNano() )
		randomInt := rand.Intn( keysLength )
		key := markov.keys[ randomInt ]

		for i := 0; i < wordLimit; i++ {
			words := strings.Split( key, ">>" )
			output += words[ 0 ] + " "

			keyPossibilities := len( markov.store[key] )
			if keyPossibilities < 1{
				// Randomly choose a possibility from the key's position in the store
				rand.Seed( time.Now().UnixNano() )
				randomInt = rand.Intn( len( markov.store[key] ) )
				key = words[1] + ">>" + markov.store[key][randomInt]
			} else{
				// If the key has 0 possibilities, randomly choose a new key to work with
				rand.Seed( time.Now().UnixNano() )
				randomInt := rand.Intn( keysLength )
				key = markov.keys[ randomInt ]
			}
		}
	}

	return strings.Replace( output, "\n", " ", -1)
}