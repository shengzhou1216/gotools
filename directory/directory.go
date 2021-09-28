package directory

import (
	"io"
	"os"
)

// Exists check if directory(dir/file) exists.
func Exists(dir string) (bool, error) {
	_, err := os.Stat(dir)
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

// RemoveDirs remove multiple dirs
func RemoveDirs(dirs ...string) error {
	for _, v := range dirs {
		if err := os.RemoveAll(v); err != nil {
			return err
		}
	}
	return nil
}

// IsEmpty check if dir is empty
func IsEmpty(dir string) (bool, error) {
	f, err := os.Open(dir)
	if err != nil {
		return false, err
	}
	defer f.Close()
	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, nil
}

// Ensure ensure directory is exists.
func Ensure(dir string) error {
	exists, err := Exists(dir)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	return nil
}
