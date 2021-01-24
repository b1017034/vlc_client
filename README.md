# vlc_go_meas

## Prerequisites

[Install Linux](https://github.com/adrg/libvlc-go/wiki/Install-on-Linux)

[Install Windows](https://github.com/adrg/libvlc-go/wiki/Install-on-Windows)

[Install Mac](https://github.com/b1017034/vlc_go_meas/wiki/Install-on-Mac)
## Installation

```
go get github.com/adrg/libvlc-go/v3
```

## RTP Stream to VLC

### MacOS

```
ffmpeg \
    -f avfoundation \
    -framerate 30 \
    -s 1280x720 \
    -pix_fmt yuyv422 \
    -re \
    -i 0:1 \
    -vf "drawtext=fontfile=/Library/Fonts/Arial.ttf:fontsize=72:fontcolor=blue:text='%{localtime}.%{eif\\:1M*t-1K*trunc(t*1K)\\:d\\:3}'" \
    -vcodec hevc_videotoolbox \
    -tag:v hevc \
    -an \
    -f rtp \
    -sdp_file path/to/video.sdp \
    "rtp://127.0.0.1:5000”
"
```

## for UDP
```
ffmpeg \
    -framerate 30.000030 \
    -f avfoundation \
    -pix_fmt uyvy422 \
    -i "0" \
    -vf "drawtext=fontfile=/Library/Fonts/Arial.ttf:fontsize=72:fontcolor=blue:text='%{localtime}.%{eif\\:1M*t-1K*trunc(t*1K)\\:d\\:3}'" \
    -vcodec hevc_videotoolbox \
    -tag:v hevc \
    -preset ultrafast \
    -tune zerolatency \
    -an \
    -f mpegts \
    "udp://192.168.1.3:5555?pkt_size=1316"
```

## VLC client

```
/Applications/VLC.app/Contents/MacOS/VLC  path/to/video.sdp
```
