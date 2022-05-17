package avcodec

/*
#include <libavcodec/avcodec.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type AVCodecParserContext struct {
	cptr *C.struct_AVCodecParserContext
}

func ParserInit(codecId int) (*AVCodecParserContext, error) {
	ret := (*C.struct_AVCodecParserContext)(C.av_parser_init(C.int(codecId)))
	if ret == (*C.struct_AVCodecParserContext)(unsafe.Pointer(uintptr(0))) {
		return nil, errors.New("parser not found")
	}

	return &AVCodecParserContext{
		cptr: ret,
	}, nil
}
