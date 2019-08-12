package commands

import (
	"fmt"
	"../structures"
	"../prompt"
	"strconv"
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
	PREPARE_TOO_LONG_STRING = 3
	PREPARE_NEGATIVE_ID = 4
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
		return statement.prepareInsert(input)
	} else if len(input) >= 6 && input[0:6] == "select" {
		statement.StatementType = STATEMENT_SELECT
		return PREPARE_SUCCESS
	} else {
		return PREPARE_UNRECOGNIZED_STATEMENT
	}

}


/**
 * Parse input line and check input line for params length
 */
func (statement *Statement) prepareInsert(input string) PrepareResult {

	statement.StatementType = STATEMENT_INSERT

	var (
		id string
		username string
		email string
	)

	scannedItems, err := fmt.Sscanf(input, "insert %s %s %s", &id, &username, &email)
	if err != nil || scannedItems < 3 {
		return PREPARE_SYNTAX_ERROR
	}

	ID, _ := strconv.Atoi(id)

	if ID < 0 {
		return PREPARE_NEGATIVE_ID
	}

	if len(username) > structures.COLUMN_USERNAME_SIZE || len(email) > structures.COLUMN_EMAIL_SIZE {
		return PREPARE_TOO_LONG_STRING
	}

	statement.Row.ID = uint32(ID)
	statement.Row.Username = username
	statement.Row.Email = email

	return PREPARE_SUCCESS

}


/**
 * Execute the prepared statement
 */
func (statement *Statement) Execute(table *structures.Table) ExecuteResult {

	switch statement.StatementType {
		case STATEMENT_INSERT:
			return statement.insertData(table)
		case STATEMENT_SELECT:
			return statement.selectData(table)
	}

	return STATEMENT_UNDEFIND

}


/**
 * Execute insert statement
 */
func (statement *Statement) insertData(table *structures.Table) ExecuteResult {

	if table.NumRows >= structures.TABLE_MAX_ROWS {
		return EXECUTE_TABLE_FULL
	}

	row := statement.Row
	//When inserting a row, we open a cursor at the end of table, write to that cursor location,
	//then close the cursor.
	cursor := table.End()

	structures.SerializeRow(&row, cursor.Value())
	table.NumRows++

	return EXECUTE_SUCCESS

}


/**
 * Execute select statement
 */
func (statement *Statement) selectData(table *structures.Table) ExecuteResult {

	//When selecting all rows in the table, we open a cursor at the start of the table,
	//print the row, then advance the cursor to the next row. Repeat until we’ve reached the end of the table.
	cursor := table.Start()
	row := structures.Row{}

	for !cursor.EndOfTable {
		structures.DeserializeRow(cursor.Value(), &row)
		prompt.PrintRow(&row)
		cursor.Advance()
	}

	return EXECUTE_SUCCESS

}