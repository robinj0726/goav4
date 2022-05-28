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

func (f *AVFrame) FramePtr() unsafe.Pointer {
	return unsafe.Pointer(f.cptr)
}

func (f *AVFrame) DataPtrPtr() unsafe.Pointer {
	return unsafe.Pointer(&f.cptr.data[0])
}

func (f *AVFrame) LineSizePtr() unsafe.Pointer {
	return unsafe.Pointer(&f.cptr.linesize[0])
}

func (f *AVFrame) DataPtr(index int) unsafe.Pointer {
	return unsafe.Pointer(&f.cptr.data[index])
}

func (f *AVFrame) LineSize(index int) int {
	return int(f.cptr.linesize[index])
}

func (f *AVFrame) Samples() int {
	return int(f.cptr.nb_samples)
}
