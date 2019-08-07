package prompt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"github.com/gookit/color"
)

var Reader *bufio.Reader
var Writer *bufio.Writer
var RedColor func(a ...interface{}) string
var YellowColor func(a ...interface{}) string

/**
 * Init BufferReader and BufferWriter
 * Init console colors
 */
func init() {

	Reader = bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	Writer = bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	// colors
	RedColor = color.FgRed.Render
	YellowColor = color.FgYellow.Render

}


/**
 * Print empty string in prompt with message to start using console
 */
func PrintPrompt() {

	fmt.Fprintf(Writer,"db > ")
	Flush()

}


/**
 * Read line from the console
 */
func ReadLine() string {
	str, _, err := Reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
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



