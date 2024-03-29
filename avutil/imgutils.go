package avutil

/*
#cgo pkg-config: libavutil
#include <libavutil/imgutils.h>
*/
import "C"
import "unsafe"

func GetImageBufferSize(pix_fmt int,
	width int32, height int32, align int) int {
	return int(C.av_image_get_buffer_size((C.enum_AVPixelFormat)(pix_fmt), (C.int)(width), (C.int)(height), (C.int)(align)))
}

func FillImageArrays(dst *AVFrame, src unsafe.Pointer,
	pix_fmt int, width int32, height int32, align int) {
	C.av_image_fill_arrays((**C.uchar)(&dst.cptr.data[0]), (*C.int)(&dst.cptr.linesize[0]), (*C.uchar)(src), (C.enum_AVPixelFormat)(pix_fmt), (C.int)(width), (C.int)(height), (C.int)(align))
}

func ImageAlloc(frame *AVFrame, width int32, height int32, pix_fmt int, align int) {
	C.av_image_alloc((**C.uchar)(&frame.cptr.data[0]), (*C.int)(&frame.cptr.linesize[0]), C.int(width), C.int(height), int32(pix_fmt), C.int(align))
}
