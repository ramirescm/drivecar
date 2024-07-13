package users

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

	h := handler{db}

	u := User{Name: "Gooper", Login: "gooper@go.com.br", Password: "123456"}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(&u)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder() // save response when request finish
	req := httptest.NewRequest(http.MethodPost, "/", &b)

	mock.
		ExpectExec(`insert into "users" ("name", "login", "password", "modified_at", "created_at")*`).
		WithArgs(u.Name, u.Login, AnyPassword{}, u.ModifiedAt, u.CreatedAt).
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

	u, err := New("Gooper", "gooper@go.com.br", "123456")
	if err != nil {
		t.Error(err)
	}

	mock.
		ExpectExec(`insert into "users" ("name", "login", "password", "modified_at", "created_at")*`).
		WithArgs("Gooper", "gooper@go.com.br", u.Password, u.ModifiedAt, u.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = Insert(db, u)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
