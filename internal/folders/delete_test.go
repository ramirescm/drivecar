package folders

import (
	"context"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi/v5"
)

func TestDeleteHTTP(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	h := handler{db}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/{id}", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	fileRows := sqlmock.
		NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, 1, 1, "maria.png", "image/png", "/", time.Now(), time.Now(), false).
		AddRow(2, 1, 1, "jose.png", "image/png", "/", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "files" where folder_id = $1 and "deleted" = false`)).
		WithArgs(1).
		WillReturnRows(fileRows)

	mock.
		ExpectExec(regexp.QuoteMeta(`update "files" set "name" = $1, "modified_at" = $2, "deleted" = $3 where id = $4`)).
		WithArgs("maria.png", AnyTime{}, true, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.
		ExpectExec(regexp.QuoteMeta(`update "files" set "name" = $1, "modified_at" = $2, "deleted" = $3 where id = $4`)).
		WithArgs("jose.png", AnyTime{}, true, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))

	folderRows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(2, 3, "travel", time.Now(), time.Now(), false).
		AddRow(4, 4, "travel", time.Now(), time.Now(), false).
		AddRow(5, 0, "pets", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "folders" where "parent_id" = $1 and "deleted" = false`)).
		WithArgs(1).
		WillReturnRows(folderRows)

	mock.
		ExpectExec(regexp.QuoteMeta(`update "folders" set "modified_at" = $1, "deleted" = true where id = $2`)).
		WithArgs(AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	h.Delete(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("Error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	mock.
		ExpectExec(regexp.QuoteMeta(`update "folders" set "modified_at" = $1, "deleted" = true where id = $2`)).
		WithArgs(AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = Delete(db, 1)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}

}
