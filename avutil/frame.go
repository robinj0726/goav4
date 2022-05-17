package avutil

/*
#include <libavutil/frame.h>
*/
import "C"

type AVFrame struct {
	cptr *C.struct_AVFrame
}

func FrameAlloc() *AVFrame {
	return &AVFrame{
		cptr: (*C.struct_AVFrame)(C.av_frame_alloc()),
	}
}

func (f *AVFrame) Free() {
	C.av_frame_free(&f.cptr)
}
