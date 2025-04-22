package storage

import (
	"os"
)

type wasmFileLock struct {
}

func (fl *wasmFileLock) release() error {
	return nil
}

func newFileLock(path string, readOnly bool) (fl fileLock, err error) {
	return &wasmFileLock{}, nil
}

func rename(oldpath, newpath string) error {
	if _, err := os.Stat(newpath); err == nil {
		if err := os.Remove(newpath); err != nil {
			return err
		}
	}

	return os.Rename(oldpath, newpath)
}

func syncDir(name string) error {
	return nil
}
