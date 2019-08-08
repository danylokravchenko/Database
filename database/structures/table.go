package structures
import "C"

const (
	PAGE_SIZE uint32 = 4096
	TABLE_MAX_PAGES = 100
	ROWS_PER_PAGE = PAGE_SIZE / ROW_SIZE
	TABLE_MAX_ROWS = ROWS_PER_PAGE * TABLE_MAX_PAGES
)

type Table struct {
	NumRows uint32
	Pages [TABLE_MAX_PAGES]*Page
}


/**
 * Free row slot
 * Return row where to write
 */
func (table *Table) RowSlot(rowNum uint32) *Row {

	pageNum := rowNum / ROWS_PER_PAGE
	page := table.Pages[pageNum]
	rowOffset := rowNum * ROWS_PER_PAGE

	return &page.Rows[rowOffset]

}


/**
 * Create new table with pages
 */
func NewTable() *Table {
	pages := [TABLE_MAX_PAGES]*Page{}
	for i := uint32(0); i < TABLE_MAX_PAGES; i++ {
		pages[i] = &Page{
			Rows: [ROWS_PER_PAGE]Row{},
		}
	}
	return &Table{
		NumRows: 0,
		Pages:   pages,
	}
}


func (table *Table) Clear() {

}