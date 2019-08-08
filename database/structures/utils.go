package structures

//#include <string.h>
import "C"
import "unsafe"


/**
 * Wrapper for C memcpy function
 */
func memcpy(destination unsafe.Pointer, source unsafe.Pointer, size uint32) {

	C.memcpy(destination, source, C.size_t(size))

}

func memmove(dest, src []byte) int {
	n := len(src)
	if len(dest) < len(src) {
		n = len(dest)
	}
	if n == 0 {
		return 0
	}
	C.memmove(unsafe.Pointer(&dest[0]), unsafe.Pointer(&src[0]), C.size_t(n))
	return n
}
