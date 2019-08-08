package commands

import (
	"fmt"
	"../structures"
	"../prompt"
)

type StatementType int

const (
	STATEMENT_INSERT StatementType = 0
	STATEMENT_SELECT = 1
	STATEMENT_UNDEFIND = 2
)

type Statement struct {
	StatementType StatementType
	Row structures.Row
}

type PrepareResult int

const (
	PREPARE_SUCCESS PrepareResult = 0
	PREPARE_UNRECOGNIZED_STATEMENT = 1
	PREPARE_SYNTAX_ERROR = 2
)

type ExecuteResult int

const (
	EXECUTE_SUCCESS ExecuteResult = 0
	EXECUTE_DUPLICATE_KEY = 1
	EXECUTE_TABLE_FULL = 2
)


/**
 * Check input command and prepare statement for execution
 */
func (statement *Statement) Prepare(input string) PrepareResult {

	if len(input) >= 6 && input[0:6] == "insert" {
		statement.StatementType = STATEMENT_INSERT
		scannedItems, err := fmt.Sscanf(input, "insert %d %s %s", &statement.Row.ID, &statement.Row.Username, &statement.Row.Email)
		if err != nil || scannedItems < 3 {
			return PREPARE_SYNTAX_ERROR
		}
		return PREPARE_SUCCESS
	} else if len(input) >= 6 && input[0:6] == "select" {
		statement.StatementType = STATEMENT_SELECT
		return PREPARE_SUCCESS
	} else {
		return PREPARE_UNRECOGNIZED_STATEMENT
	}

}


/**
 * Execute the prepared statement
 */
func (statement *Statement) Execute(table *structures.Table) ExecuteResult {

	switch statement.StatementType {
		case STATEMENT_INSERT:
			return statement.Insert(table)
		case STATEMENT_SELECT:
			return statement.Select(table)
	}

	return STATEMENT_UNDEFIND

}


/**
 * Execute insert statement
 */
func (statement *Statement) Insert(table *structures.Table) ExecuteResult {

	if table.NumRows >= structures.TABLE_MAX_ROWS {
		return EXECUTE_TABLE_FULL
	}

	row := statement.Row

	structures.SerializeRow(&row, table.RowSlot(table.NumRows))
	table.NumRows++

	return EXECUTE_SUCCESS

}


/**
 * Execute select statement
 */
func (statement *Statement) Select(table *structures.Table) ExecuteResult {

	row := structures.Row{}

	for i := uint32(0) ; i < table.NumRows; i++ {
		structures.DeserializeRow(table.RowSlot(i), &row)
		prompt.PrintRow(&row)
	}

	return EXECUTE_SUCCESS

}