package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestCreate() {
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(ts.entity)
	assert.NoError(ts.T(), err)

	ts.entity.SetPassword(ts.entity.Password)

	rr := httptest.NewRecorder() // save response when request finish
	req := httptest.NewRequest(http.MethodPost, "/", &b)

	setMockInsert(ts.mock, ts.entity)

	ts.handler.Create(rr, req)

	assert.Equal(ts.T(), http.StatusCreated, rr.Code)
}

func (ts *TransactionSuite) TestInsert() {
	setMockInsert(ts.mock, ts.entity)
	_, err := Insert(ts.conn, ts.entity)
	assert.NoError(ts.T(), err)
}

func setMockInsert(mock sqlmock.Sqlmock, entity *User) {
	mock.
		ExpectExec(`insert into "users" ("name", "login", "password", "modified_at", "created_at")*`).
		WithArgs(entity.Name, entity.Login, entity.Password, entity.ModifiedAt, entity.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
