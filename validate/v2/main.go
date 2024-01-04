package main

import (
	cds "github.com/craterdog/go-cdsn-validation/v3"
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
	var parser = cds.ParserClass().Default()
	var document = parser.ParseDocument(string(bytes))

	// Validate the parse tree.
	var validator = cds.ValidatorClass().Default()
	validator.ValidateDocument(document)

	// Format the grammar file.
	var formatter = cds.FormatterClass().Default()
	var formatted = formatter.FormatDocument(document)

	// Write out the formatted file.
	err = osx.WriteFile(filename, []byte(formatted), 0644)
	if err != nil {
		panic(err)
	}
}
