package check

import (
	"path/filepath"
)

type FileCheck interface {
	Run(string) error
	RunAll([]string) error
}

type FileOptions struct {
	BasePath string
}

func (opts *FileOptions) FullPath(path string) string {
	if opts.BasePath != "" {
		return filepath.Join(opts.BasePath, path)
	}

	return path
}
