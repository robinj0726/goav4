package avutil

/*
#include <libavutil/frame.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type AVFrame struct {
	cptr *C.struct_AVFrame
}

func (f AVFrame) String() string {
	return fmt.Sprintf("%#v", *(f.cptr))
}

func FrameAlloc() *AVFrame {
	return &AVFrame{
		cptr: (*C.struct_AVFrame)(C.av_frame_alloc()),
	}
}

func (f *AVFrame) Free() {
	C.av_frame_free(&f.cptr)
}

func (f *AVFrame) FrameRef() unsafe.Pointer {
	return unsafe.Pointer(f.cptr)
}
