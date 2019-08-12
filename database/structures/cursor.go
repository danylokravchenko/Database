package structures

type Cursor struct {
	Table *Table
	RowNumber uint32
	EndOfTable bool // Indicates a position one past the last element
}


/**
 * Free row slot
 * Return row where to write
 */
func (cursor *Cursor) Value() *Row {

	rowNum := cursor.RowNumber
	pageNum := rowNum / ROWS_PER_PAGE
	page := cursor.Table.Pager.GetPage(pageNum)
	rowOffset := rowNum % ROWS_PER_PAGE

	return &page.Rows[rowOffset]

}


/**
 * Move cursor to the next row
 */
func (cursor *Cursor) Advance() {

	cursor.RowNumber++

	if cursor.RowNumber >= cursor.Table.NumRows {
		cursor.EndOfTable = true
	}

}