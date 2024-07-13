package folders

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	f, err := New("photos", 0)
	if err != nil {
		t.Error(err)
	}

	f.ModifiedAt = time.Now()
	mock.
		ExpectExec(regexp.QuoteMeta(`insert into "folders" ("parent_id", "name", "modified_at") values ($1, $2, $3)`)).
		WithArgs(0, "photos", f.ModifiedAt).
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
