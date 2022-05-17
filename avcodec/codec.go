package avcodec

/*
#include <libavcodec/codec.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type AVCodec struct {
	cptr *C.struct_AVCodec
}

func FindDecoder(codecId int) (*AVCodec, error) {
	ret := (*C.struct_AVCodec)(C.avcodec_find_decoder(C.enum_AVCodecID(codecId)))
	if ret == (*C.struct_AVCodec)(unsafe.Pointer(uintptr(0))) {
		return nil, errors.New("Unsupported codec!")
	}
	return &AVCodec{
		cptr: ret,
	}, nil
}
