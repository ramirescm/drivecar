package files

import (
	"errors"
	"time"
)

var (
	ErrorOwnerRequired = errors.New("Owner is required")
	ErrorNameRequired  = errors.New("Name is required")
	ErrorTypeRequired  = errors.New("Type is required")
	ErrorPathRequired  = errors.New("Path is required")
)

type File struct {
	ID         int64     `json:"id"`
	FolderID   int64     `json:"-"`
	OwnerID    int64     `json:"owner_id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Path       string    `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
}

func (f *File) Validate() error {
	if f.OwnerID == 0 {
		return ErrorOwnerRequired
	}

	if f.Name == "" {
		return ErrorNameRequired
	}

	if f.Type == "" {
		return ErrorTypeRequired
	}

	if f.Path == "" {
		return ErrorPathRequired
	}

	return nil
}
