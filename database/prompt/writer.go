package prompt

import (
	"bufio"
	"fmt"
	"os"
)

var Writer *bufio.Writer

func init() {

	Writer = bufio.NewWriterSize(os.Stdout, 1024 * 1024)

}



/**
 * Print empty string in prompt with message to start using console
 */
func PrintPrompt() {

	fmt.Fprintf(Writer,"db > ")
	Flush()

}


/**
 * Flush writer's buffer
 */
func Flush() {
	Writer.Flush()
}


/**
 * Print error message in console using red color
 */
func PrintError(message string) {

	fmt.Fprintf(Writer,"%s \n", RedColor(message))
	Flush()

}


func PrintInfoMessage(message string) {

	fmt.Fprintf(Writer, "%s \n", YellowColor(message))
	Flush()
}