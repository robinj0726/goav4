package avformat

/*
#cgo pkg-config: libavformat libavcodec libavutil
#include <libavformat/avformat.h>

struct buffer_data {
    uint8_t *ptr;
    size_t size; ///< size left in the buffer
};

static int read_packet(void *opaque, uint8_t *buf, int buf_size)
{
    struct buffer_data *bd = (struct buffer_data *)opaque;
    buf_size = FFMIN(buf_size, bd->size);
    if (!buf_size)
        return AVERROR_EOF;
    printf("ptr:%p size:%zu\n", bd->ptr, bd->size);
    return buf_size;
}

AVFormatContext* allocate_avformat_context()
{
	AVFormatContext *fmt_ctx = NULL;
	AVIOContext *avio_ctx = NULL;
	uint8_t *buffer = NULL, *avio_ctx_buffer = NULL;
    size_t buffer_size, avio_ctx_buffer_size = 4096;

	struct buffer_data bd = { 0 };
	bd.ptr  = buffer;
    bd.size = buffer_size;

	if (!(fmt_ctx = avformat_alloc_context())) {
		return NULL;
	}

	avio_ctx_buffer = av_malloc(avio_ctx_buffer_size);
    if (!avio_ctx_buffer) {
		return NULL;
    }
    avio_ctx = avio_alloc_context(avio_ctx_buffer, avio_ctx_buffer_size,
                                  0, &read_packet, NULL, NULL, NULL);
    if (!avio_ctx) {
		return NULL;
	}
    fmt_ctx->pb = avio_ctx;
	return fmt_ctx;
}
*/
import "C"
import "unsafe"

type AVFormatContext struct {
	Ctx *C.struct_AVFormatContext
}

func AllocContext() *AVFormatContext {
	return &AVFormatContext{
		Ctx: (*C.struct_AVFormatContext)(C.allocate_avformat_context()),
	}
}

func (c *AVFormatContext) FreeContext() {
	C.avformat_free_context((*C.struct_AVFormatContext)(unsafe.Pointer(c.Ctx)))
}

func Version() uint {
	return uint(C.avformat_version())
}
