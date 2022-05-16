package avformat

/*
#cgo pkg-config: libavformat libavcodec libavutil
#include <libavformat/avformat.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type AVFormatContext struct {
	ctx *C.struct_AVFormatContext
}

func AllocContext() *AVFormatContext {
	return &AVFormatContext{
		ctx: (*C.struct_AVFormatContext)(C.avformat_alloc_context()),
	}
}

func (c *AVFormatContext) FreeContext() {
	C.avformat_free_context((*C.struct_AVFormatContext)(unsafe.Pointer(c.ctx)))
}

func (c *AVFormatContext) OpenInput(url string) error {
	ret := (int)(C.avformat_open_input((**C.struct_AVFormatContext)(&c.ctx), C.CString(url), (*C.struct_AVInputFormat)(unsafe.Pointer(uintptr(0))), (**C.struct_AVDictionary)(unsafe.Pointer(uintptr(0)))))
	if ret < 0 {
		return fmt.Errorf("Could not open source file %s", url)
	}
	return nil
}

func (c *AVFormatContext) CloseInput() {
	C.avformat_close_input((**C.struct_AVFormatContext)(&c.ctx))
}

func (c *AVFormatContext) FindStreamInfo() error {
	ret := (int)(C.avformat_find_stream_info((*C.struct_AVFormatContext)(c.ctx), (**C.struct_AVDictionary)(unsafe.Pointer(uintptr(0)))))
	if ret < 0 {
		return errors.New("Could not find stream information")
	}
	return nil
}

func (c *AVFormatContext) DumpFormat() {
	C.av_dump_format((*C.struct_AVFormatContext)(c.ctx), 0, C.CString(""), 0)
}

func Version() uint {
	return uint(C.avformat_version())
}
