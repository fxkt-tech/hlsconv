package hls

import (
	"context"
	"os"

	"github.com/fxkt-tech/hlsconv/internal/conf"
	"github.com/fxkt-tech/liv/ffmpeg"
	"github.com/fxkt-tech/liv/ffmpeg/codec"
	"github.com/fxkt-tech/liv/ffmpeg/filter"
	"github.com/fxkt-tech/liv/ffmpeg/input"
	"github.com/fxkt-tech/liv/ffmpeg/naming"
	"github.com/fxkt-tech/liv/ffmpeg/output"
)

func TranscodeHls(infile, outfolder string, c *conf.Config) error {
	os.MkdirAll(outfolder, os.ModePerm)

	var (
		ctx = context.Background()
		nm  = naming.New()

		input1 = input.WithSimple(infile)

		scale1 = filter.Scale(nm.Gen(), -2, -2).Use(filter.SelectStream(0, filter.StreamVideo, true))
	)

	return ffmpeg.NewFFmpeg(
		ffmpeg.Binary(c.Codec.FFmpegCmd),
		ffmpeg.V(ffmpeg.LogLevelError),
		ffmpeg.Debug(true),
		// ffmpeg.Dry(true),
	).AddInput(
		input1,
	).AddFilter(
		scale1,
	).AddOutput(
		output.New(
			output.Map(scale1.Name(0)),
			output.Map("0:a?"),
			output.VideoCodec(codec.X264),
			output.AudioCodec(codec.Copy),
			output.File(outfolder+"m.m3u8"),
			output.MovFlags("faststart"),
			output.HLSSegmentType("mpegts"),
			output.HLSFlags("independent_segments"),
			output.HLSPlaylistType("vod"),
			output.HLSTime(c.Codec.HlsTime),
			output.HLSKeyInfoFile(c.Codec.HlsKeyinfoFile), // 加密
			output.HLSSegmentFilename(outfolder+"m-%5d.ts"),
			output.Format(codec.HLS),
		),
	).Run(ctx)
}
