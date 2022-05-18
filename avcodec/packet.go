package avcodec

/*
#include <libavcodec/avcodec.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type AVPacket struct {
	cptr *C.struct_AVPacket
}

func PacketAlloc() *AVPacket {
	return &AVPacket{
		cptr: (*C.struct_AVPacket)(C.av_packet_alloc()),
	}
}

func (pkt *AVPacket) Free() {
	C.av_packet_free(&pkt.cptr)
}

func (pkt *AVPacket) PacketRef() unsafe.Pointer {
	return unsafe.Pointer(pkt.cptr)
}

func (pkt AVPacket) String() string {
	return fmt.Sprintf("%#v", *(pkt.cptr))
}
