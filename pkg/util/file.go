package util

import (
	"io"
	"os"
	"path/filepath"

	multierror "github.com/hashicorp/go-multierror"
)

func CopyFile(source string, dests ...string) error {
	// Open the source file.
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	length := len(dests)
	ws := make([]io.Writer, length)
	dsts := make([]*os.File, length)
	closeDsts := func(end int) error {
		var result error
		for i := 0; i < end; i++ {
			if err := dsts[i].Close(); err != nil {
				result = multierror.Append(result, err)
			}
		}
		return result
	}

	for i, dest := range dests {
		// Makes the directory needed to create the dst
		// file.
		err = os.MkdirAll(filepath.Dir(dest), 0775)
		if err != nil {
			closeDsts(i)
			return err
		}

		// Create the destination file.
		dst, err := os.Create(dest)
		if err != nil {
			closeDsts(i)
			return err
		}

		ws[i] = dst
		dsts[i] = dst
	}

	defer closeDsts(length)

	// Copy the contents of the file.
	_, err = io.Copy(io.MultiWriter(ws...), src)
	if err != nil {
		return err
	}

	// Copy the mode if the user can't
	// open the file.
	info, err := os.Stat(source)
	if err != nil {
		var result error
		for _, dest := range dests {
			if err := os.Chmod(dest, info.Mode()); err != nil {
				result = multierror.Append(result, err)
			}
		}
		if result != nil {
			return result
		}
	}

	return nil
}
