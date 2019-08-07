package main

import (
	"log"
	"./prompt"
	"./errors"
	"os"
)

func main() {

	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	for {

		prompt.PrintPrompt()

		str := prompt.ReadLine()

		if str == "" {
			errors.EmptyLine()
		}

		if (str == "exit") {
			prompt.PrintInfoMessage("Bye")
			os.Exit(0)
		} else {
			errors.UnrecognizedCommand(str)
		}

	}

}