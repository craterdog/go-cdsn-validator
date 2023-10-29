package main

import (
	cds "github.com/craterdog/go-cdsn-validation/v2"
	osx "os"
)

// MAIN PROGRAM

func main() {
	// Validate the commandline arguments.
	if len(osx.Args) < 2 {
		panic("validate <grammar file>")
	}
	var filename = osx.Args[1]

	// Read in the grammar file.
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Parse and validate the grammar file.
	var document = cds.ParseDocument(bytes)

	// Format the grammar file.
	bytes = cds.FormatDocument(document)

	// Write out the formatted file.
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
