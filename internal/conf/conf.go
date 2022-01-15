package conf

import (
	"io/ioutil"

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
	fbytes, err := ioutil.ReadFile(conf)
	if err != nil {
		return nil, err
	}

	var bc Config
	err = yaml.Unmarshal(fbytes, &bc)
	if err != nil {
		return nil, err
	}

	return &bc, nil
}
