package gotools

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
)

// Md5Bytes md5 encode bytes
func Md5Bytes(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

// Md5Str md5 encode string
func Md5Str(s string) string {
	return Md5Bytes([]byte(s))
}

// Md5FilePath md5 encode file
func Md5FilePath(fullpath string) (string, error) {
	f, err := os.Open(fullpath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return Md5File(f)
}

// Md5File md5 encode file
func Md5File(f *os.File) (r string, err error) {
	h := md5.New()
	if _, err = io.Copy(h, f); err != nil {
		return
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// Md5MultipartFile md5 encode multipart file
func Md5MultipartFile(f *multipart.File) (r string, err error) {
	h := md5.New()
	if _, err = io.Copy(h, *f); err != nil {
		return
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
