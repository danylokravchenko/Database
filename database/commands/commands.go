package commands

import "../structures"

type MetaCommandResult int

const (
	META_COMMAND_SUCCESS MetaCommandResult = 0
	META_COMMAND_UNRECOGNIZED_COMMAND = 1
	META_COMMAND_EXIT = 2
)

/**
 * Check input command and prepare constant for switch statement
 */
func MetaCommand(input string, table *structures.Table) MetaCommandResult {

	if len(input) >= 5 && input[0:5] == ".exit" {
		table.DBClose()
		return META_COMMAND_EXIT
	} else {
		return META_COMMAND_UNRECOGNIZED_COMMAND
	}

}
