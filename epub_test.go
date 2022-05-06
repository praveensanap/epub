package epub_test

import (
	"github.com/praveensanap/epub"
	"testing"
)

func TestEpub(t *testing.T) {
	bk, err := open(t, "test.epub")
	if err != nil {
		t.Fatal(err)
	}
	defer bk.Close()
}

func open(t *testing.T, f string) (*epub.Book, error) {
	bk, err := epub.Open(f)
	if err != nil {
		return nil, err
	}
	defer bk.Close()

	t.Logf("files: %+v", bk.Files())
	t.Logf("book: %+v", bk)

	return bk, nil
}
