package avcodec

/*
#cgo pkg-config: libavcodec
#include <libavcodec/avcodec.h>
*/
import "C"

type AVCodecContext struct {
	cptr *C.struct_AVCodecContext
}
