package downloader

import (
	"context"
	"io"

	"github.com/vbauerster/mpb/v4"
	"github.com/vbauerster/mpb/v4/decor"
)

type State struct {
	Bar           *mpb.Bar
	LocalFilePath string
	Results       []string
}

type Step interface {
	GetProgressParams() (int64, decor.Decorator)
	Execute(ctx context.Context, s *State) error
}

type StepWriter interface {
	GetWriter() (io.Writer, error)
}

type DownloadAsset interface {
	GetName() string
	GetSize() int64
	Open(ctx context.Context) (io.ReadCloser, error)
}

type StepProvider interface {
	GetDownloaderSteps(ctx context.Context) ([]Step, error)
}
