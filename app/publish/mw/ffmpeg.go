package mw

import (
	"bytes"
	"context"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
	"google.golang.org/grpc/metadata"
)

// GetSnapshot get the first frame of a video via ffmpeg
func GetSnapshot(videoPath string) (buf *bytes.Buffer, err error) {
	buf = bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf).
		Run()

	return buf, nil
}

// URLconvert Convert the path in the database into a complete url accessible by the front end
func URLconvert(ctx context.Context, path string) string {
	if len(path) == 0 {
		return ""
	}
	arr := strings.Split(path, "/")
	u, err := GetObjURL(ctx, arr[0], arr[1])
	if err != nil {
		return ""
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	u.Scheme = md["scheme"][0]
	u.Host = md["host"][0]
	u.Path = "/src" + md["path"][0]
	return u.String()
}
