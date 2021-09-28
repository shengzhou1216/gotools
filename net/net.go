package net

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)



// DownloadFile download file to dest. file's parent directory should be pre-created.
func DownloadFile(url string, dest string) error {
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if _, err := io.Copy(f, resp.Body); err != nil {
		return err
	}
	// status code >= 200 && status code < 300
	if !(resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return errors.New(resp.Status)
	}
	return nil
}

// UploadMultipartFile upload multipart file.
func UploadMultipartFile(url string, f *multipart.FileHeader) (*http.Response, error) {
	file, err := f.Open()
	if err != nil {
		return nil, err
	}
	return Upload(url, file, f.Filename)
}

// UploadFile upload file.
func UploadFile(url string, f *os.File) (*http.Response, error) {
	return Upload(url, f, f.Name())
}

// UploadPath upload file by path.
func UploadPath(url string, fullpath string) (resp *http.Response, err error) {
	f, err := os.Open(fullpath)
	if err != nil {
		return
	}
	return Upload(url, f, path.Base(fullpath))
}

// Upload file
func Upload(url string, reader io.Reader, name string) (resp *http.Response, err error) {
	body := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", name)
	if err != nil {
		return
	}
	if _, err = io.Copy(part, reader); err != nil {
		return
	}
	_ = writer.Close()
	resp, err = http.Post(url, writer.FormDataContentType(), body)
	if err != nil {
		return
	}
	_ = resp.Body.Close()
	return
}
