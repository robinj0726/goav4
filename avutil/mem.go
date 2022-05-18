package avutil

/*
#include <libavutil/avutil.h>
*/
import "C"
import "unsafe"

func Malloc(size int) unsafe.Pointer {
	return unsafe.Pointer(C.av_malloc(C.size_t(size)))
}

func Free(ptr unsafe.Pointer) {
	C.av_free(ptr)
}
