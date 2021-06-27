package common

import "os"

func (f String) IsDirectory() bool {
	if file, err := os.Stat(f.Trim()); err != nil {
		return os.IsExist(err)
	} else {
		return file.IsDir()
	}
}

func (f String) IsFileExist() bool {
	_, err := os.Stat(f.Trim())
	return err == nil || os.IsExist(err)
}

func (f String) Open() (*os.File, error) {
	return os.Open(f.Trim())
}
