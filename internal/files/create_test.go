package files

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ramirescm/drivecar/internal/bucket"
	"github.com/ramirescm/drivecar/internal/queue"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	b, err := bucket.New(bucket.MockProvider, nil)
	if err != nil {
		t.Error(err)
	}

	q, err := queue.New(queue.MockProvider, nil)
	if err != nil {
		t.Error(err)
	}

	h := handler{db, b, q}

	// upload
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	file, err := os.Open("./testdata/test.jpg")
	if err != nil {
		t.Error(err)
	}

	w, err := mw.CreateFormFile("file", "test.jpg")
	if err != nil {
		t.Error(err)
	}

	_, err = io.Copy(w, file)
	if err != nil {
		t.Error(err)
	}
	mw.Close()
	// end upload

	rr := httptest.NewRecorder() // save response when request finish
	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Add("Content-Type", mw.FormDataContentType())

	mock.
		ExpectExec(`insert into "files" ("folder_id", "owner_id", "name", "type", "path", "modified_at")*`).
		WithArgs(0, 1, "test.jpg", "application/octet-stream", "/test.jpg", AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	h.Create(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	f, err := New(1, "Gooper.png", "image/png", "/")
	if err != nil {
		t.Error(err)
	}

	mock.
		ExpectExec(`insert into "files" ("folder_id", "owner_id", "name", "type", "path", "modified_at")*`).
		WithArgs(f.FolderID, f.OwnerID, f.Name, f.Type, f.Path, f.ModifiedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = Insert(db, f)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
