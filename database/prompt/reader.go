package prompt

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var Reader *bufio.Reader

func init() {

	Reader = bufio.NewReaderSize(os.Stdin, 1024 * 1024)

}


/**
 * Read line from the console and trim a line
 */
func ReadLine() string {

	str, _, err := Reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimLeft(strings.TrimRight(string(str), "\r\n"), " ")

}