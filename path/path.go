package path

import (
	"io"
	"os"
)

// Exists check if path(dir/file) exists.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDirs create multiple dir
func CreateDirs(dirs ...string) error {
	for _, v := range dirs {
		exist, err := Exists(v)
		if err != nil {
			return err
		}
		if !exist {
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				return err
			}
		}
	}
	return nil
}

// IsEmptyDir check if dir is empty
func IsEmptyDir(dir string) (bool, error) {
	f, err := os.Open(dir)
	if err != nil {
		return false, err
	}
	defer f.Close()
	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return false, err
	}
	return true, nil
}

// EnsurePath ensure path is exists.
func EnsurePath(p string) error {
	exists, err := Exists(p)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	if err := os.MkdirAll(p, os.ModePerm); err != nil {
		return err
	}
	return nil
}
