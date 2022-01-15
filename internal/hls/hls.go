package hls

import (
	"context"
	"os"

	"fxkt.tech/ffmpeg"
	"fxkt.tech/ffmpeg/codec"
	"fxkt.tech/ffmpeg/filter"
	"fxkt.tech/ffmpeg/input"
	"fxkt.tech/ffmpeg/output"
	"fxkt.tech/hlsconv/internal/conf"
)

func TranscodeHls(infile, outfolder string, c *conf.Config) error {
	os.MkdirAll(outfolder, os.ModePerm)
	ff := ffmpeg.Default()
	ff.CmdLoc(c.Codec.FFmpegCmd)
	ff.LogLevel("error")
	ff.AddInput(input.New(
		input.I(infile),
	))
	ff.AddOutput(output.New(
		output.VideoCoder(codec.VideoX264),
		output.AudioCoder(codec.Copy),
		output.Map(filter.SelectStream(0, filter.StreamVideo, true)),
		output.Map(filter.SelectStream(0, filter.StreamAudio, false)),
		output.File(outfolder+"m.m3u8"),
		output.MovFlags("faststart"),
		output.HlsSegmentType("mpegts"),
		output.HlsFlags("independent_segments"),
		output.HlsPlaylistType("vod"),
		output.HlsTime(c.Codec.HlsTime),
		output.HlsKeyInfoFile(c.Codec.HlsKeyinfoFile), // 加密
		output.HlsSegmentFilename(outfolder+"m-%5d.ts"),
		output.Format(codec.Hls),
	))
	return ff.Run(context.Background())
}
