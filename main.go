package main

import (
	"fmt"
	"log"

	"github.com/fxkt-tech/hlsconv/internal/conf"
	"github.com/fxkt-tech/hlsconv/internal/errors"
	"github.com/fxkt-tech/hlsconv/internal/hls"
	"github.com/fxkt-tech/hlsconv/pkg/file"
	"github.com/fxkt-tech/hlsconv/pkg/flag"
)

func main() {
	// 解析命令行参数
	args := flag.ParseFlags()
	// 判断-i参数是否是以/结尾来判断处理单个文件还是批量处理
	cfg, err := conf.Init(args.ConfFile)
	if err != nil {
		errors.Blask(err)
	}
	var files []string
	if file.IsFolder(args.Input) {
		// 如果是批量处理，则遍历该文件夹下所有以mp4结尾的文件
		var err error
		files, err = file.ParseFilesBySuffix(args.Input, ".mp4", args.Recursion)
		if err != nil {
			errors.Blask(err)
		}
	} else {
		files = append(files, args.Input)
	}
	// 处理所有转码任务
	for i, infile := range files {
		outfolder := fmt.Sprintf("%s%s/", args.Output, file.FileName(infile))
		err := hls.TranscodeHls(infile, outfolder, cfg)
		status := "succ"
		if err != nil {
			status = fmt.Sprintf("fail[%v]", err)
		}
		log.Printf("#%d %s, infile: %s, outfolder: %s\n", i, status, infile, outfolder)
	}
}
