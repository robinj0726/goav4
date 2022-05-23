package avutil

/*
#include <libavutil/avutil.h>
*/
import "C"

const (
	AV_PIX_FMT_NONE    = C.AV_PIX_FMT_NONE
	AV_PIX_FMT_YUV420P = C.AV_PIX_FMT_YUV420P ///< planar YUV 4:2:0, 12bpp, (1 Cr & Cb sample per 2x2 Y samples)
	AV_PIX_FMT_YUYV422 = C.AV_PIX_FMT_YUYV422 ///< packed YUV 4:2:2, 16bpp, Y0 Cb Y1 Cr
	AV_PIX_FMT_RGB24   = C.AV_PIX_FMT_RGB24   ///< packed RGB 8:8:8, 24bpp, RGBRGB...
	AV_PIX_FMT_BGR24   = C.AV_PIX_FMT_BGR24   ///< packed RGB 8:8:8, 24bpp, BGRBGR...
	AV_PIX_FMT_YUV422P = C.AV_PIX_FMT_YUV422P ///< planar YUV 4:2:2, 16bpp, (1 Cr & Cb sample per 2x1 Y samples)
	AV_PIX_FMT_YUV444P = C.AV_PIX_FMT_YUV444P ///< planar YUV 4:4:4, 24bpp, (1 Cr & Cb sample per 1x1 Y samples)
	AV_PIX_FMT_YUV410P = C.AV_PIX_FMT_YUV410P ///< planar YUV 4:1:0,  9bpp, (1 Cr & Cb sample per 4x4 Y samples)
	AV_PIX_FMT_YUV411P = C.AV_PIX_FMT_YUV411P ///< planar YUV 4:1:1, 12bpp, (1 Cr & Cb sample per 4x1 Y samples)

)
