package structures

const (
	PAGE_SIZE uint32 = 4096
	TABLE_MAX_PAGES = 100
	ROWS_PER_PAGE = PAGE_SIZE / ROW_SIZE
	TABLE_MAX_ROWS = ROWS_PER_PAGE * TABLE_MAX_PAGES
	FILE_LENGTH = 154
	ROW_LENGTH = 12
)

type Table struct {
	NumRows uint32
	Pager *Pager
}


/**
 * Open connection to db and initialize table and pager
 */
func DBOpen(filename string) *Table {

	pager := OpenFile(filename)
	numRows := (pager.FileLength - FILE_LENGTH) / ROW_LENGTH//pager.FileLength / ROW_SIZE

	return &Table{
		NumRows: numRows,
		Pager:   pager,
	}

}


/**
 * Flush the page cache to disk, close the database file, free the memory for the Pager and Table data structures
 */
func (table *Table) DBClose() {

	pager := table.Pager

	defer pager.FileDescriptor.Close()

	numFullPages := table.NumRows / ROWS_PER_PAGE

	for i := uint32(0); i < numFullPages; i++ {
		if pager.Pages[i] == nil {
			continue
		}
		pager.Flush(i, PAGE_SIZE)
		pager.Pages[i] = nil
	}

	// There may be a partial page to write to the end of the file
	// This should not be needed after we switch to a B-tree
	numAdditionalRows := table.NumRows % ROWS_PER_PAGE

	if numAdditionalRows > 0 {
		pageNum := numFullPages
		if pager.Pages[pageNum] != nil {
			pager.Flush(pageNum, numAdditionalRows*ROW_SIZE)
			pager.Pages[pageNum] = nil
		}
	}

	for i := uint32(0); i < TABLE_MAX_PAGES; i++ {
		pager.Pages[i] = nil
	}

}


/**
 * Create new cursor at the start of the table
 */
func (table *Table) Start() *Cursor {

	return &Cursor{
		Table:      table,
		RowNumber:  0,
		EndOfTable: table.NumRows == 0,
	}

}


/**
 * Create new cursor at the end of the table
 */
func (table *Table) End() *Cursor {

	return &Cursor{
		Table:      table,
		RowNumber:  table.NumRows,
		EndOfTable: true,
	}

}