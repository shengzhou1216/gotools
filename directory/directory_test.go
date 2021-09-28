package directory

import "testing"

func TestDirectory(t *testing.T) {
	t.Run("test dir exists", func(t *testing.T) {
		exists, err := Exists("directory.go")
		if err != nil {
			t.Fatal(err)
		}
		t.Log("is test dir exists?", exists)
	})

	t.Run("test create dirs", func(t *testing.T) {
		err := CreateDirs("test1", "test2", "test3/test1/test2")
		if err != nil {
			t.Fatal(err)
		}

		//RemoveDirs("test1", "test2", "test3")
	})

	t.Run("test check if dir is empty", func(t *testing.T) {
		r, err := IsEmpty("test")
		if err != nil {
			t.Fatal(err)
		}
		t.Log("is dir test is empty?", r)
	})

}
