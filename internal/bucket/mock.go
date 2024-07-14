package bucket

import (
	"io"
	"os"
)

type MockBucket struct {
	content map[string][]byte
}

func (mb *MockBucket) Upload(file io.Reader, key string) error {
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	mb.content[key] = data
	return nil
}

func (mb *MockBucket) Download(src, dst string) (*os.File, error) {
	panic("unimplemented")
}

func (mb *MockBucket) Delete(key string) error {
	return nil
}
