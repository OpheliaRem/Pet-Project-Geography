package logger

import (
	"os"
)

var fStreams []*os.File

func AddFileStream(file *os.File) {
	fStreams = append(fStreams, file)
}

func Log(m string) {
	for _, stream := range fStreams {
		stream.Write([]byte(m + "\n"))
	}
}
