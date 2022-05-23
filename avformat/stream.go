package avformat

/*
#include <libavformat/avformat.h>
int av_stream_codec_type(AVStream* stream)
{
	return stream->codecpar->codec_type;
}

int av_stream_codec_id(AVStream* stream)
{
	return stream->codecpar->codec_id;
}
*/
import "C"
import (
	"unsafe"
)

type AVStream struct {
	cptr *C.struct_AVStream
}

func (s *AVStream) CodecType() int {
	return int(C.av_stream_codec_type(s.cptr))
}

func (s *AVStream) CodecID() int {
	return int(C.av_stream_codec_id(s.cptr))
}

func (s *AVStream) CodecParameters() unsafe.Pointer {
	return unsafe.Pointer(s.cptr.codecpar)
}
