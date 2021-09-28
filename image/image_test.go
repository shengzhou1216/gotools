package image

import "testing"

func TestImage(t *testing.T) {
	t.Run("test base64 encode", func(t *testing.T) {
		r, err := Base64Image("yby.jpg")
		if err != nil {
			t.Fatal(err)
		}
		t.Log("base64", r)
	})

}
