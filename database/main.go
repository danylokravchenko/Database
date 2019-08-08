package main

import (
	"log"
	"os"
	"./prompt"
	"./errors"
	"./commands"
)

func main() {

	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

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
		switch commands.PrepareStatement(input, statement) {
			case commands.PREPARE_SUCCESS:
				break
			case commands.PREPARE_UNRECOGNIZED_STATEMENT:
				errors.UnrecognizedKeyword(input)
				continue
		}
		commands.ExecuteStatement(statement)
		prompt.PrintInfoMessage("Executed")

	}

}