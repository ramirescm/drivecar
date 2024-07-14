package files

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	h := handler{db, nil, nil}

	f := File{Name: "Gooper"}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(&f)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder() // save response when request finish
	req := httptest.NewRequest(http.MethodPost, "/", &b)

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
