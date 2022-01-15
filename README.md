# hlsconv
hls converter.

## 外部依赖
> ffmpeg: 4.2.1

## 使用方式
```bash
linix/macos
# 转换单个文件
bin/hlsconv -i video/in.mp4 -o outvideo/
# 批量转换文件
bin/hlsconv -i video/ -o outvideo/

windows
# 转换单个文件
.\bin\hlsconv.exe -i .\video\in.mp4 -o .\outvideo\
# 批量转换文件
.\bin\hlsconv.exe -i .\video\ -o .\outvideo\
```

转码后的hls分片文件会放在<kbd>-o</kbd>所在目录下和源文件同名的文件夹下

## 转码配置
在conf/app.yaml中可以控制部分转码配置
```yaml
codec:
  hls_time: 2 # 分片时长
  hls_keyinfo_file: lib/file.keyinfo # hls加密信息文件
```

## 加密设置
在config.yaml中指定<kbd>hls_keyinfo_file</kbd>参数即可开启hls加密；若开启加密，则在之前需要生成file.key，然后将file.key放置于一个公网可访问的位置（比如云存储），然后替换file.keyinfo文件的第一行。
在lib/file.keyinfo中给了基本案例，在实际处理中需要先修改本文件。