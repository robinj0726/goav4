package avcodec

/*
#cgo pkg-config: libavcodec
#include <libavcodec/avcodec.h>
*/
import "C"
import (
	"errors"
	"fmt"
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

func (ctx AVCodecContext) String() string {
	return fmt.Sprintf("%#v", *(ctx.cptr))
}

func (ctx *AVCodecContext) Open2(codec *AVCodec) error {
	ret := C.avcodec_open2((*C.struct_AVCodecContext)(ctx.cptr), (*C.struct_AVCodec)(codec.cptr), (**C.struct_AVDictionary)(unsafe.Pointer(uintptr(0))))
	if int(ret) < 0 {
		return errors.New("Could not open codec")
	}
	return nil
}

func (ctx *AVCodecContext) ParametersToContext(par unsafe.Pointer) {
	C.avcodec_parameters_to_context(ctx.cptr, (*C.struct_AVCodecParameters)(par))
}

func (ctx *AVCodecContext) Width() int32 {
	return int32(ctx.cptr.width)
}

func (ctx *AVCodecContext) Height() int32 {
	return int32(ctx.cptr.height)
}

func (ctx *AVCodecContext) PixFmt() int {
	return int(ctx.cptr.pix_fmt)
}

func (ctx *AVCodecContext) SampleRate() int {
	return int(ctx.cptr.sample_rate)
}

func (ctx *AVCodecContext) Channels() int {
	return int(ctx.cptr.channels)
}

func (ctx *AVCodecContext) SampleFmt() int {
	return int(ctx.cptr.sample_fmt)
}

func (ctx *AVCodecContext) CPtr() unsafe.Pointer {
	return unsafe.Pointer(ctx.cptr)
}

func (ctx *AVCodecContext) SendPacket(pktref unsafe.Pointer) error {
	ret := C.avcodec_send_packet(ctx.cptr, (*C.struct_AVPacket)(pktref))
	if ret < 0 {
		return errors.New("send packet to codec error")
	}

	return nil
}

func (ctx *AVCodecContext) ReceiveFrame(frameref unsafe.Pointer) error {
	ret := C.avcodec_receive_frame(ctx.cptr, (*C.struct_AVFrame)(frameref))
	if ret < 0 {
		return errors.New("receive frame from codec error")
	}

	return nil
}

func NewContextWithCPtr(cptr unsafe.Pointer) *AVCodecContext {
	return &AVCodecContext{
		cptr: (*C.struct_AVCodecContext)(cptr),
	}
}
