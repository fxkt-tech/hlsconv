package conf

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Codec *Codec `yaml:"codec"`
}

type Codec struct {
	FFmpegCmd      string `yaml:"ffmpeg_cmd"`
	HlsTime        int32  `yaml:"hls_time"`
	HlsKeyinfoFile string `yaml:"hls_keyinfo_file"`
}

func Init(conf string) (*Config, error) {
	f, err := os.Open(conf)
	if err != nil {
		return nil, err
	}

	bs, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var bc Config
	err = yaml.Unmarshal(bs, &bc)
	if err != nil {
		return nil, err
	}

	return &bc, nil
}
