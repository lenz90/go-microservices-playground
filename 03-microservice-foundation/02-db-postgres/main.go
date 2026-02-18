package main

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"net/http"
	"os"
)

type User struct{ ID, Name, Email string }

func main() {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/playground?sslmode=disable"
	}
	db, _ := sql.Open("pgx", dsn)
	defer db.Close()
	_, _ = db.Exec(`CREATE TABLE IF NOT EXISTS users(id TEXT PRIMARY KEY,name TEXT,email TEXT)`)
	r := chi.NewRouter()
	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		var u User
		_ = json.NewDecoder(r.Body).Decode(&u)
		_, _ = db.Exec(`INSERT INTO users(id,name,email) VALUES($1,$2,$3)`, u.ID, u.Name, u.Email)
		w.WriteHeader(201)
	})
	r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(User{ID: "id"}) })
	_ = http.ListenAndServe(":8080", r)
}
