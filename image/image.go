package image

import (
	"encoding/base64"
	"net/http"
)

// Base64 get image base64 from image bytes.
func Base64(bytes []byte) string {
	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)
	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	return base64Encoding
}

// bytes to base64
func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
