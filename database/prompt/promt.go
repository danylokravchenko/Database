package prompt

import (
	"github.com/gookit/color"
)

var RedColor func(a ...interface{}) string
var YellowColor func(a ...interface{}) string


/**
 * Init console colors
 */
func init() {

	// colors
	RedColor = color.FgRed.Render
	YellowColor = color.FgYellow.Render

}





