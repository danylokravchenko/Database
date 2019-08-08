package structures

import "unsafe"

const (
	COLUMN_USERNAME_SIZE = 32
	COLUMN_EMAIL_SIZE = 225
)

const (
	ID_SIZE uint32 = uint32(unsafe.Sizeof(new(Row).ID))
	USERNAME_SIZE uint32 = uint32(unsafe.Sizeof(new(Row).Username))
	EMAIL_SIZE uint32 = uint32(unsafe.Sizeof(new(Row).Email))
	ID_OFFSET = 0
	USERNAME_OFFSET = ID_OFFSET + ID_SIZE
	EMAIL_OFFSET = USERNAME_OFFSET + USERNAME_SIZE
	ROW_SIZE = ID_SIZE + USERNAME_SIZE + EMAIL_SIZE
)

type Row struct {
	ID uint32
	Username string//[COLUMN_USERNAME_SIZE]rune
	Email string//[COLUMN_EMAIL_SIZE]rune
}


/**
 * Copy bytes from source Row to memory
 */
func SerializeRow(source *Row, destination *Row) {

	destination.ID = source.ID
	destination.Username = source.Username
	destination.Email = source.Email

	//memcpy(unsafe.Pointer(&destination.ID), unsafe.Pointer(&source.ID), ID_SIZE)
	//memcpy(unsafe.Pointer(&destination.Username), unsafe.Pointer(&source.Username), USERNAME_SIZE)
	//memcpy(unsafe.Pointer(&destination.Email), unsafe.Pointer( &source.Email), EMAIL_SIZE)

	//destination.ID = source.ID
	//memcpy(&destination[0 + ID_OFFSET],  &source.ID, ID_SIZE)
	//memcpy(&destination[0 + USERNAME_OFFSET],  &source.Username, USERNAME_SIZE)
	//memcpy(&destination[0 + EMAIL_OFFSET],  &source.Email, EMAIL_SIZE)

}


/**
 * Copy bytes from memory to destination Row
 */
func DeserializeRow(source *Row, destination *Row) {

	destination.ID = source.ID
	destination.Username = source.Username
	destination.Email = source.Email

	//memcpy(unsafe.Pointer(&destination.ID), unsafe.Pointer(&source.ID), ID_SIZE)
	//memcpy(unsafe.Pointer(&destination.Username), unsafe.Pointer(&source.Username), USERNAME_SIZE)
	//memcpy(unsafe.Pointer(&destination.Email), unsafe.Pointer( &source.Email), EMAIL_SIZE)

	//memcpy(&destination.ID, &source[0 + ID_OFFSET], ID_SIZE)
	//memcpy(&destination.Username, &source[0 + USERNAME_OFFSET], USERNAME_SIZE)
	//memcpy(&destination.Email, &source[0 + EMAIL_OFFSET], EMAIL_SIZE)

}