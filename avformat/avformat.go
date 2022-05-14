package avformat

/*
#cgo pkg-config: libavformat libavcodec libavutil
#include <libavformat/avformat.h>
*/
import "C"
import "unsafe"

type AVFormatContext struct {
	ctx *C.struct_AVFormatContext
}

func AllocContext() *AVFormatContext {
	return &AVFormatContext{
		ctx: (*C.struct_AVFormatContext)(C.avformat_alloc_context()),
	}
}

func (f *AVFormatContext) FreeContext() {
	C.avformat_free_context((*C.struct_AVFormatContext)(unsafe.Pointer(f.ctx)))
}

func Version() uint {
	return uint(C.avformat_version())
}
