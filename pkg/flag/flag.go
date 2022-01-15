package flag

import (
	syserrors "errors"
	"flag"

	"fxkt.tech/hlsconv/internal/errors"
	"fxkt.tech/hlsconv/pkg/file"
)

type Flags struct {
	Input     string //input
	Output    string // output
	ConfFile  string // config file path.
	Recursion bool   // deal folder by rectusion.
}

func ParseFlags() Flags {
	var i, o, conf string
	var r bool
	flag.StringVar(&i, "i", "", "input file or folder.")
	flag.StringVar(&o, "o", "", "output folder.")
	flag.StringVar(&conf, "conf", "conf/config.yaml", "config file path.")
	flag.BoolVar(&r, "r", false, "deal folder by rectusion.")
	flag.Parse()
	i = file.FixWinPath(i)
	o = file.FixWinPath(o)
	conf = file.FixWinPath(conf)
	if i == "" {
		errors.Blask(syserrors.New("infile required."))
	}
	if o == "" {
		errors.Blask(syserrors.New("outfolder required."))
	}
	if !file.IsFolder(o) {
		errors.Blask(syserrors.New("outfolder must be a folder."))
	}
	return Flags{
		Input:     i,
		Output:    o,
		ConfFile:  conf,
		Recursion: r,
	}
}
