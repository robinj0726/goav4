package avcodec

/*
#cgo pkg-config: libavcodec
#include <libavcodec/avcodec.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type AVCodecContext struct {
	cptr *C.struct_AVCodecContext
}

func AllocContext3(codec *AVCodec) (*AVCodecContext, error) {
	ret := (*C.struct_AVCodecContext)(C.avcodec_alloc_context3((*C.struct_AVCodec)(codec.cptr)))
	if ret == (*C.struct_AVCodecContext)(unsafe.Pointer(uintptr(0))) {
		return nil, errors.New("Could not allocate codec context")
	}
	return &AVCodecContext{
		cptr: ret,
	}, nil

}

func (avctx *AVCodecContext) Open2(codec *AVCodec) error {
	ret := C.avcodec_open2((*C.struct_AVCodecContext)(avctx.cptr), (*C.struct_AVCodec)(codec.cptr), (**C.struct_AVDictionary)(unsafe.Pointer(uintptr(0))))
	if int(ret) < 0 {
		return errors.New("Could not open codec")
	}
	return nil
}
