package swscale

/*
#cgo pkg-config: libswscale
#include <libswscale/swscale.h>
*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/robinj730/goav4/avutil"
)

const (
	SWS_FAST_BILINEAR = 1
	SWS_BILINEAR      = 2
	SWS_BICUBIC       = 4
)

type SwsContext struct {
	cptr *C.struct_SwsContext
}

func GetContext(srcW int32, srcH int32, srcFormat int,
	dstW int32, dstH int32, dstFormat int,
	flags int) (*SwsContext, error) {

	ret := (*C.struct_SwsContext)(C.sws_getContext((C.int)(srcW), (C.int)(srcH), (int32)(srcFormat), (C.int)(dstW), (C.int)(dstH), (int32)(dstFormat), (C.int)(flags),
		(*C.struct_SwsFilter)(unsafe.Pointer(uintptr(0))), (*C.struct_SwsFilter)(unsafe.Pointer(uintptr(0))), (*C.double)(unsafe.Pointer(uintptr(0.0)))))

	if ret == (*C.struct_SwsContext)(unsafe.Pointer(uintptr(0))) {
		return nil, errors.New("Could not get sws context")
	}
	return &SwsContext{
		cptr: ret,
	}, nil
}

func (sws *SwsContext) Scale(srcFrame avutil.AVFrame, srcSliceY int32, srcSliceH int32, dstFrame avutil.AVFrame) {
	C.sws_scale((*C.struct_SwsContext)(sws.cptr), (**C.uchar)(srcFrame.Data()), (*C.int)(srcFrame.LineSize()), (C.int)(srcSliceY), (C.int)(srcSliceH), (**C.uchar)(dstFrame.Data()), (*C.int)(dstFrame.LineSize()))
}
