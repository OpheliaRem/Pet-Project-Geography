package logger

import (
	"fmt"
	"os"
)

var fStreams []*os.File

func AddFileStream(file *os.File) {
	fStreams = append(fStreams, file)
}

func Log(m any) {
	for _, stream := range fStreams {
		fmt.Fprintln(stream, m)
	}
}
