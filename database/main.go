package main

import (
	"./commands"
	"./errors"
	"./prompt"
	"./structures"
	"log"
	"os"
)

func main() {

	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	table := structures.NewTable()

	for {

		prompt.PrintPrompt()

		input := prompt.ReadLine()

		if input == "" {
			errors.EmptyLine()
		}


		if string(input[0]) == "." {
			switch commands.MetaCommand(input) {
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