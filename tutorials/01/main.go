package main

/*
#include <libavutil/frame.h>

void SaveFrame(AVFrame *pFrame, int width, int height, int iFrame) {
  FILE *pFile;
  char szFilename[32];
  int  y;

  // Open file
  sprintf(szFilename, "frame%d.ppm", iFrame);
  pFile=fopen(szFilename, "wb");
  if(pFile==NULL)
    return;

  // Write header
  fprintf(pFile, "P6\n%d %d\n255\n", width, height);

  // Write pixel data
  for(y=0; y<height; y++)
    fwrite(pFrame->data[0]+y*pFrame->linesize[0], 1, width*3, pFile);

  // Close file
  fclose(pFile);
}
*/
import "C"
import (
	"github.com/robinj730/goav4/avcodec"
	"github.com/robinj730/goav4/avformat"
	"github.com/robinj730/goav4/avutil"
	"github.com/robinj730/goav4/swscale"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Please provide a movie file")
	// 	return
	// }

	infile := "../sample.mp4"

	pFormatCtx := avformat.AVFormatContext{}
	pFormatCtx.OpenInput(infile)
	defer pFormatCtx.CloseInput()

	pFormatCtx.FindStreamInfo()

	var pCodec *avcodec.AVCodec
	var pCodecCtx *avcodec.AVCodecContext
	videoStream := -1
	i := 0
	for stream := range pFormatCtx.AVStreams() {
		if stream.CodecType() == avutil.AVMEDIA_TYPE_VIDEO {
			pCodec, _ = avcodec.FindDecoder(stream.CodecID())
			pCodecCtx, _ = avcodec.AllocContext3(pCodec)
			pCodecCtx.ParametersToContext(stream.CodecParameters())
			videoStream = i
			i += 1
		}
	}

	if videoStream == -1 {
		panic("no video stream inside")
	}

	err := pCodecCtx.Open2(pCodec)
	if err != nil {
		panic(err)
	}

	pFrame := avutil.FrameAlloc()
	defer pFrame.Free()

	pFrameRGB := avutil.FrameAlloc()
	defer pFrameRGB.Free()

	numBytes := avutil.GetImageBufferSize(int(avutil.AV_PIX_FMT_RGB24), pCodecCtx.Width(), pCodecCtx.Height(), 16)
	buffer := avutil.Malloc(numBytes)
	defer avutil.Free(buffer)

	avutil.FillImageArrays(pFrameRGB, buffer, int(avutil.AV_PIX_FMT_RGB24), pCodecCtx.Width(), pCodecCtx.Height(), 16)

	// fmt.Println(pCodecCtx)
	sws_ctx, _ := swscale.GetContext(pCodecCtx.Width(), pCodecCtx.Height(), pCodecCtx.PixFmt(), pCodecCtx.Width(), pCodecCtx.Height(), int(avutil.AV_PIX_FMT_RGB24), swscale.SWS_BILINEAR)

	pkt := avcodec.PacketAlloc()
	defer pkt.Free()

	n := 0
	for {
		err := pFormatCtx.ReadFrame(pkt.PacketRef())
		if err != nil {
			break
		}

		if pkt.StreamIndex() == videoStream {
			err := pCodecCtx.SendPacket(pkt.PacketRef())
			if err != nil {
				panic(err)
			}

			err = pCodecCtx.ReceiveFrame(pFrame.FrameRef())
			if err != nil {
				panic(err)
			}

			sws_ctx.Scale(*pFrame, 0, pCodecCtx.Height(), *pFrameRGB)
			C.SaveFrame((*C.struct_AVFrame)(pFrameRGB.FrameRef()), (C.int)(pCodecCtx.Width()), (C.int)(pCodecCtx.Height()), (C.int)(n))
			n += 1
		}

	}
}
