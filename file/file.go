package file

import (
	"io"
	"os"
)

// Copy file from source to dest.
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

// Write  bytes to file
func Write(b []byte, dest string) error {
	f, err := os.Open(dest)
	if err != nil {
		return nil
	}
	if _, err := f.Write(b); err != nil {
		return err
	}
	return nil
}

// WriteString write string to file
func WriteString(s string, dest string) error {
	f, err := os.Open(dest)
	if err != nil {
		return nil
	}
	if _, err := f.WriteString(s); err != nil {
		return err
	}
	return nil
}
