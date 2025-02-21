package log

import (
	"io"
	"log"
	"log/slog"
	"os"
	"sync"

	"github.com/hashicorp/terraform/logging"
)

var logger *slog.Logger
var once sync.Once

func SetLogger(l *slog.Logger) {
	logger = l
}

// Init initializes the log, it can only be called once,
// repetitive calls to it will not change it.
// It also set the level of vebosity of the
// Terraform logs via tflogs, if true it'll
// use the TF_LOG env variable to set it to
// Terraform
func Init(out io.Writer, tflogs bool) {
	once.Do(func() {
		level := slog.LevelInfo
		if tflogs {
			level = slog.LevelDebug
		}

		logHandler := slog.NewTextHandler(out, &slog.HandlerOptions{
			Level: level,
		})

		if !tflogs {
			os.Setenv("TF_LOG", "")

			out := logging.LogOutput()
			if out == nil {
				out = io.Discard
			}
			log.SetOutput(out)
		}

		SetLogger(slog.New(logHandler))
	})
}

// Get returns the initialized logger,
// if it has not been initialized it'll
// initialize it with the default values
func Get() *slog.Logger {
	Init(io.Discard, false)
	return logger
}
