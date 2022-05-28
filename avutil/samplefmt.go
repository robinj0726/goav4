package avutil

/*
#include <libavutil/samplefmt.h>
*/
import "C"
import "unsafe"

const (
	AV_SAMPLE_FMT_NONE = C.AV_SAMPLE_FMT_NONE
	AV_SAMPLE_FMT_U8   = C.AV_SAMPLE_FMT_U8   ///< unsigned 8 bits
	AV_SAMPLE_FMT_S16  = C.AV_SAMPLE_FMT_S16  ///< signed 16 bits
	AV_SAMPLE_FMT_S32  = C.AV_SAMPLE_FMT_S32  ///< signed 32 bits
	AV_SAMPLE_FMT_FLT  = C.AV_SAMPLE_FMT_FLT  ///< float
	AV_SAMPLE_FMT_DBL  = C.AV_SAMPLE_FMT_DBL  ///< double
	AV_SAMPLE_FMT_U8P  = C.AV_SAMPLE_FMT_U8P  ///< unsigned 8 bits, planar
	AV_SAMPLE_FMT_S16P = C.AV_SAMPLE_FMT_S16P ///< signed 16 bits, planar
	AV_SAMPLE_FMT_S32P = C.AV_SAMPLE_FMT_S32P ///< signed 32 bits, planar
	AV_SAMPLE_FMT_FLTP = C.AV_SAMPLE_FMT_FLTP ///< float, planar
	AV_SAMPLE_FMT_DBLP = C.AV_SAMPLE_FMT_DBLP ///< double, planar
	AV_SAMPLE_FMT_S64  = C.AV_SAMPLE_FMT_S64  ///< signed 64 bits
	AV_SAMPLE_FMT_S64P = C.AV_SAMPLE_FMT_S64P ///< signed 64 bits, planar
	AV_SAMPLE_FMT_NB   = C.AV_SAMPLE_FMT_NB   ///< Number of sample formats. DO NOT USE if linking dynamically
)

func GetSamplesBufferSize(linesize *int, nb_channels int, nb_samples int, sample_fmt int, align int) int {
	return int(C.av_samples_get_buffer_size((*C.int)(unsafe.Pointer(linesize)), C.int(nb_channels), C.int(nb_samples), int32(sample_fmt), C.int(align)))
}
