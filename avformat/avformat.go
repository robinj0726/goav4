package avformat

/*
#cgo pkg-config: libavformat libavcodec libavutil
#include <libavformat/avformat.h>

AVStream* av_get_stream(AVFormatContext* fmt_ctx, int index) {
    return fmt_ctx->streams[index];
}
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

func Version() uint {
	return uint(C.avformat_version())
}

type AVFormatContext struct {
	cptr *C.struct_AVFormatContext
}

func AllocContext() *AVFormatContext {
	return &AVFormatContext{
		cptr: (*C.struct_AVFormatContext)(C.avformat_alloc_context()),
	}
}

func (ctx *AVFormatContext) FreeContext() {
	C.avformat_free_context((*C.struct_AVFormatContext)(unsafe.Pointer(ctx.cptr)))
}

func (ctx *AVFormatContext) OpenInput(url string) error {
	ret := (int)(C.avformat_open_input((**C.struct_AVFormatContext)(&ctx.cptr), C.CString(url), (*C.struct_AVInputFormat)(unsafe.Pointer(uintptr(0))), (**C.struct_AVDictionary)(unsafe.Pointer(uintptr(0)))))
	if ret < 0 {
		return fmt.Errorf("Could not open source file %s", url)
	}
	return nil
}

func (ctx *AVFormatContext) CloseInput() {
	C.avformat_close_input((**C.struct_AVFormatContext)(&ctx.cptr))
}

func (ctx *AVFormatContext) FindStreamInfo() error {
	ret := (int)(C.avformat_find_stream_info((*C.struct_AVFormatContext)(ctx.cptr), (**C.struct_AVDictionary)(unsafe.Pointer(uintptr(0)))))
	if ret < 0 {
		return errors.New("Could not find stream information")
	}
	return nil
}

func (ctx *AVFormatContext) DumpFormat() {
	C.av_dump_format((*C.struct_AVFormatContext)(ctx.cptr), 0, C.CString(""), 0)
}

func (ctx *AVFormatContext) AVStreams() <-chan AVStream {
	ch := make(chan AVStream)

	go func() {
		for i := 0; i < int(ctx.cptr.nb_streams); i++ {
			ch <- AVStream{
				cptr: C.av_get_stream(ctx.cptr, (C.int)(i)),
			}
		}
		close(ch)
	}()

	return ch
}

func (ctx *AVFormatContext) ReadFrame(pktref unsafe.Pointer) error {
	ret := C.av_read_frame(ctx.cptr, (*C.struct_AVPacket)(pktref))
	if ret < 0 {
		return errors.New("av_read_frame error")
	}

	return nil
}
