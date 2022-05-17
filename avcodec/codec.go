package avcodec

import "C"

type AVCodec struct {
	cptr *C.struct_AVCodec
}
