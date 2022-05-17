package avcodec

/*
#include <libavcodec/avcodec.h>
*/
import "C"

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
