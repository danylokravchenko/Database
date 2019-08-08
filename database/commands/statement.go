package commands

import "fmt"

type StatementType int

const (
	STATEMENT_INSERT StatementType = 0
	STATEMENT_SELECT = 1
)

type Statement struct {
	statement_type StatementType
}

type PrepareResult int

const (
	PREPARE_SUCCESS PrepareResult = 0
	PREPARE_UNRECOGNIZED_STATEMENT = 1
)

/**
 * Check input command and prepare statement for execution
 */
func PrepareStatement(input string, statement *Statement) PrepareResult {

	if input == "insert" {
		statement.statement_type = STATEMENT_INSERT
		return PREPARE_SUCCESS
	} else if input == "select" {
		statement.statement_type = STATEMENT_SELECT
		return PREPARE_SUCCESS
	} else {
		return PREPARE_UNRECOGNIZED_STATEMENT
	}

}


/**
 * Execute the prepared statement
 */
func ExecuteStatement(statement *Statement) {

	switch statement.statement_type {
		case STATEMENT_INSERT:
			fmt.Println("This is where we would do an insert.")
			break
		case STATEMENT_SELECT:
			fmt.Println("This is where we would do a select.")
			break
	}
	
}