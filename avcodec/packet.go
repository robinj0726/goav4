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

func (pkt AVPacket) String() string {
	return fmt.Sprintf("%#v", *(pkt.cptr))
}

func PacketAlloc() *AVPacket {
	return &AVPacket{
		cptr: (*C.struct_AVPacket)(C.av_packet_alloc()),
	}
}

func (pkt *AVPacket) Free() {
	C.av_packet_free(&pkt.cptr)
}

func (pkt *AVPacket) PacketPtr() unsafe.Pointer {
	return unsafe.Pointer(pkt.cptr)
}

func (pkt *AVPacket) StreamIndex() int {
	return int(pkt.cptr.stream_index)
}

func (pkt *AVPacket) PacketClone() *AVPacket {
	return &AVPacket{
		cptr: C.av_packet_clone(pkt.cptr),
	}
}
