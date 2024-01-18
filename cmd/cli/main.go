package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/timdevlet/mp4/internal/configs"
	"github.com/timdevlet/mp4/internal/logs"
	"github.com/timdevlet/mp4/internal/mp4"
)

func main() {
	opt := configs.NewConfigsFromEnv()
	logs.InitLog(opt.LOG_FORMAT, opt.LOG_LEVEL)

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide a file path")
		os.Exit(1)
	}

	path := args[0:]

	for _, p := range path {
		res, err := mp4Info(p)

		if err != nil {
			logrus.Error(err)
		} else {
			for _, r := range res {
				fmt.Println(r)
			}
		}

		fmt.Println()
	}
}

func mp4Info(fn string) (res []string, err error) {
	f, err := os.Open(fn)
	if err != nil {
		return res, err
	}
	defer f.Close()

	println("file:", fn)

	media, err := mp4.Parse(f)
	if err != nil {
		return res, err
	}

	for _, c := range media.GetTracks() {
		w, h, err := c.WidthAndHeight()
		if err == nil && w > 0 && h > 0 {
			res = append(res, fmt.Sprintf("video: %v %v", w, h))
			continue
		}

		bitrate, err := c.AudioBitrate()
		if err == nil && bitrate > 1000 {
			res = append(res, fmt.Sprintf("audio: %v hz", bitrate))
		}
	}

	return res, err
}
