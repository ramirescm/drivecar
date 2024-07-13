package folders

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	db *sql.DB
}

func SetRoutes(r chi.Router, db *sql.DB) {
	h := handler{db}

	r.Post("", h.Create)
	r.Put("/{id}", h.Update)
	r.Get("/{id}", h.Get)
	r.Delete("/{id}", h.Delete)
}
