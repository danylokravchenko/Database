package errors

import (
	"../prompt"
	"fmt"
)


/**
 * Error for entering an empty line by user
 */
func EmptyLine() {
	prompt.PrintError("Please, enter a command!")
}


/**
 * Error for command that not present
 */
func UnrecognizedCommand(message string) {
	prompt.PrintError(fmt.Sprintf("Unrecognized command '%s'.", message))
}
