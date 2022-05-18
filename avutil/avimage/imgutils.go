package avimage

/*
#cgo pkg-config: libavutil
#include <libavutil/imgutils.h>
*/
import "C"

func GetBufferSize(pix_fmt int,
	width int, height int, align int) int {
	return int(C.av_image_get_buffer_size((C.enum_AVPixelFormat)(pix_fmt), (C.int)(width), (C.int)(height), (C.int)(align)))
}
