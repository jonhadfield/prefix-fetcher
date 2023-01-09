package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func saveFile(i saveFileInput) error {
	pf, err := os.Stat(i.path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	if pf != nil && pf.IsDir() {
		i.path = filepath.Join(i.path, i.defaultFileName)
	}

	f, err := os.Create(i.path)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.Write(i.data); err != nil {
		return err
	}

	fmt.Printf("%s data written to: %s", i.provider, filepath.Clean(i.path))

	return nil
}

type saveFileInput struct {
	provider        string
	data            []byte
	path            string
	defaultFileName string
}
