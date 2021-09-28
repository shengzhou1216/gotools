package net

import "testing"

func TestNet(t *testing.T) {
	t.Run("test download file", func(t *testing.T) {
		url := "https://img9.doubanio.com/view/subject/l/public/s9114855.jpg"
		err := DownloadFile(url, "test.png")
		if err != nil {
			t.Fatal(err)
		}
	})

}
