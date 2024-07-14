package files

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi/v5"
)

func TestModify(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	h := handler{db, nil, nil}

	f := File{ID: 1, Name: "maria.png"}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(&f)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder() // save response when request finish
	req := httptest.NewRequest(http.MethodPut, "/{id}", &b)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	rows := sqlmock.
		NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, 1, 1, "maria.png", "image/png", "/", time.Now(), time.Now(), false)

	mock.
		ExpectQuery(regexp.QuoteMeta(`select * from "files" where id = $1`)).
		WithArgs(1).
		WillReturnRows(rows)

	mock.
		ExpectExec(regexp.QuoteMeta(`update "files" set "name" = $1, "modified_at" = $2, "deleted" = $3 where id = $4`)).
		WithArgs(f.Name, AnyTime{}, false, f.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	h.Modify(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	mock.
		ExpectExec(regexp.QuoteMeta(`update "files" set "name" = $1, "modified_at" = $2, "deleted" = $3 where id = $4`)).
		WithArgs("Batman", AnyTime{}, false, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = Update(db, 1, &File{Name: "Batman"})
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
