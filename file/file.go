package file

import (
	"github.com/shengzhou1216/gotools/directory"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// Copy file from source to dest. dest file directory should be pre-created.
func Copy(source, dest string) error {
	s, err := os.Open(source)
	if err != nil {
		return err
	}
	defer s.Close()
	//var d *os.File
	d, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer d.Close()
	if _, err := io.Copy(d, s); err != nil {
		return err
	}
	return nil
}

// Write  bytes to file. if dest's parent directory is not exists, it will be auto crated.
func Write(b []byte, dest string) error {
	// ensure parent dir exists
	if err := directory.Ensure(path.Dir(dest)); err != nil {
		return err
	}
	f, err := os.Create(dest)
	if err != nil {
		return nil
	}
	if _, err := f.Write(b); err != nil {
		return err
	}
	return nil
}

// WriteString write string to dest. if dest's parent directory is not exists, it will be auto crated.
// if dest is exists,it's content will be replaced by string s.
func WriteString(s string, dest string) error {
	// ensure parent dir exists
	if err := directory.Ensure(path.Dir(dest)); err != nil {
		return err
	}
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(s); err != nil {
		return err
	}
	return nil
}

// ReadBytes read file as bytes
func ReadBytes(dest string) ([]byte, error) {
	return ioutil.ReadFile(dest)
}

// ReadString read file content as string
func ReadString(dest string) (string, error) {
	bytes, err := ioutil.ReadFile(dest)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

