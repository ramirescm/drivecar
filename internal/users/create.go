package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
	u := new(User)

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	u.SetPassword(u.Password)

	err = u.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := Insert(h.db, u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	u.ID = id

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(u)
}

func Insert(db *sql.DB, u *User) (id int64, err error) {
	stmt := `insert into "users" ("name", "login", "password", "modified_at", "created_at") VAUES($1, $2, $3, $4, $5)`
	result, err := db.Exec(stmt, u.Name, u.Login, u.Password, u.ModifiedAt, u.CreatedAt)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
