package file

import "testing"

func TestFile(t *testing.T) {
	t.Run("test copy file", func(t *testing.T) {
		err := Copy("test.md", "copy.md")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("test write string to file", func(t *testing.T) {
		err := WriteString("hello world", "hello.md")
		if err != nil {
			t.Fatal(err)
		}
	})
}
