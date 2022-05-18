package avcodec

/*
#include <libavcodec/codec.h>
*/
import "C"
import (
	"errors"
	"fmt"
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

func (a AVCodec) String() string {
	return fmt.Sprintf("%#v", *(a.cptr))
}

func (a AVCodec) Name() string {
	return fmt.Sprintf("%s", C.GoString(a.cptr.name))
}

func (a AVCodec) LongName() string {
	return fmt.Sprintf("%s", C.GoString(a.cptr.long_name))
}
