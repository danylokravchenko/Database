package structures

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

type Page struct {
	Rows [ROWS_PER_PAGE]Row
}


// Page struct in file to get off empty rows with slice of them
type DBPage struct {
	Rows []Row
}

/**
 * init gob encoder
 */
func init() {

	gob.Register(Page{})
	gob.Register(DBPage{})

}

/**
 * Copy not empty rows to DBPage struct
 */
func (page *Page) ToDBPage() DBPage {

	var rows []Row

	for i := uint32(0); i < ROWS_PER_PAGE; i++ {
		if page.Rows[i].Username != "" && page.Rows[i].Email != "" {
			rows = append(rows, page.Rows[i])
		}
	}

	return DBPage{
		Rows:rows,
	}

}


/**
 * Copy not empty rows from DBPage to Page struct
 */
func (dbpage DBPage) GetPage() *Page {

	page := Page{
		Rows: [ROWS_PER_PAGE]Row{},
	}

	for i := 0; i < len(dbpage.Rows); i++ {
		page.Rows[i] = dbpage.Rows[i]
	}

	return &page

}


/**
 * Go binary encoder
 */
func (page DBPage) ToGOB64() string {

	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(page)
	if err != nil { fmt.Println(`failed gob Encode`, err) }

	return base64.StdEncoding.EncodeToString(b.Bytes())

}


/**
 * Go binary decoder
 */
func FromGOB64(str string) *DBPage {

	page := DBPage{}
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil { fmt.Println(`failed base64 Decode`, err); }
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&page)
	if err != nil { fmt.Println(`failed gob Decode`, err); }

	return &page

}