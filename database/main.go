package main

import (
	"./commands"
	"./errors"
	"./prompt"
	"./structures"
	"log"
	"os"
	"flag"
)

func main() {

	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var filename string
	flag.StringVar(&filename, "file", "", "Usage")
	flag.Parse()

	if filename == "" {
		errors.DatabaseFilename()
		os.Exit(0)
	}

	table := structures.DBOpen(filename)

	for {

		prompt.PrintPrompt()

		input := prompt.ReadLine()

		if input == "" {
			errors.EmptyLine()
		}


		if string(input[0]) == "." {
			switch commands.MetaCommand(input, table) {
				case commands.META_COMMAND_SUCCESS:
					continue
				case commands.META_COMMAND_UNRECOGNIZED_COMMAND:
					errors.UnrecognizedCommand(input)
					continue
				case commands.META_COMMAND_EXIT:
					prompt.PrintInfoMessage("Bye")
					os.Exit(0)
			}
		}

		statement := &commands.Statement{}
		switch statement.Prepare(input) {
			case commands.PREPARE_SUCCESS:
				break
			case commands.PREPARE_SYNTAX_ERROR:
				errors.SyntaxError()
				continue
			case commands.PREPARE_UNRECOGNIZED_STATEMENT:
				errors.UnrecognizedKeyword(input)
				continue
			case commands.PREPARE_TOO_LONG_STRING:
				errors.StringTooLong()
				continue
			case commands.PREPARE_NEGATIVE_ID:
				errors.NegativeId()
				continue
		}

		switch statement.Execute(table) {
			case commands.EXECUTE_SUCCESS:
				prompt.PrintInfoMessage("Executed")
				break
			case commands.EXECUTE_TABLE_FULL:
				errors.TableFull()
				break
		}


	}

}