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


/**
 * Error for keyword that not present
 */
func UnrecognizedKeyword(message string) {
	prompt.PrintError(fmt.Sprintf("Unrecognized keyword at start of '%s'.", message))
}


/**
 * Error for syntax error
 */
func SyntaxError() {
	prompt.PrintError("Syntax error. Could not execute the statement.")
}


/**
 * Error for full table
 */
func TableFull() {
	prompt.PrintError("Error: Table full.")
}