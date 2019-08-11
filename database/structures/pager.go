package structures

import (
	"fmt"
	"io"
	"os"
)

type Pager struct {
	FileDescriptor *os.File
	FileLength uint32
	Pages [TABLE_MAX_PAGES]*Page
}


/**
 * Open connection to a file and initialize new Pager
 */
func OpenFile(filename string) *Pager {

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666) //0755
	if err != nil {
		fmt.Printf("Unable to create file %v \n", err)
		os.Exit(1)
	}

	fileLength, err := file.Seek(0, io.SeekEnd)

	pages := [TABLE_MAX_PAGES]*Page{}
	for i := uint32(0); i < TABLE_MAX_PAGES; i++ {
		// because now db is using filesystem so we need to check if page is cached
		pages[i] = nil
		//pages[i] = &Page{
		//	Rows: [ROWS_PER_PAGE]Row{},
		//}
	}

	return &Pager{
		FileDescriptor: file,
		FileLength:     uint32(fileLength),
		Pages:          pages,
	}

}


func (pager *Pager) GetPage(pageNumber uint32) *Page {

	if pageNumber > TABLE_MAX_PAGES {
		fmt.Println(fmt.Sprintf("Tried to fetch page number out of bounds. %d > %d\n", pageNumber, TABLE_MAX_PAGES))
		os.Exit(1)
	}

	var page *Page

	if pager.Pages[pageNumber] == nil {
		// Cache miss. Allocate memory and load from file.
		numPages := pager.FileLength / PAGE_SIZE

		// We might save a partial page at the end of the file
		// TODO: test it more, maybe it's not "!=" , but "=="
		if pager.FileLength % PAGE_SIZE != 0 {
			numPages++
		}

		if pageNumber <= numPages {
			pager.FileDescriptor.Seek(int64(pageNumber * PAGE_SIZE), io.SeekStart)

			data := make([]byte, PAGE_SIZE)

			for{
				_, err := pager.FileDescriptor.Read(data)
				if err != nil {
					if err == io.EOF{   // end of the file
						if string(data) == string(make([]byte, PAGE_SIZE)) {
							page = &Page{
								Rows: [ROWS_PER_PAGE]Row{},
							}
						}
						break
					} else {
						fmt.Println("Error reading file");
						os.Exit(1)
					}
				}
				dbPage := FromGOB64(string(data))
				page = dbPage.GetPage()
			}
		}

		pager.Pages[pageNumber] = page

	}

	return pager.Pages[pageNumber]

}


/**
 * Write data to a file
 */
func (pager *Pager) Flush(pageNum uint32, size uint32) {

	if pager.Pages[pageNum] == nil {
		fmt.Println("Tried to flush nil page")
		os.Exit(1)
	}

	_, err := pager.FileDescriptor.Seek(int64(pageNum * PAGE_SIZE), io.SeekStart)
	if err != nil {
		fmt.Println("Error seeking")
		os.Exit(1)
	}

	_, err = pager.FileDescriptor.Write([]byte(pager.Pages[pageNum].ToDBPage().ToGOB64()))
	if err != nil {
		fmt.Println("Error writing")
		os.Exit(1)
	}

}


/**
 * Count not empty rows
 */
func (pager *Pager) CountRows() uint32 {

	var counter uint32

	for i := uint32(0); i < TABLE_MAX_PAGES; i++ {
		for j := uint32(0); j < ROWS_PER_PAGE; j++ {
			if pager.Pages[i] != nil && &pager.Pages[i].Rows[j] != new(Row) {
				counter++
			}
		}
	}

	return counter

}