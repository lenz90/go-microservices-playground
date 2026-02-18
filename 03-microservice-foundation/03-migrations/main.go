package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
	"path/filepath"
	"sort"
)

func apply(db *sql.DB, dir string) error {
	_, _ = db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations(filename TEXT PRIMARY KEY)`)
	es, _ := os.ReadDir(dir)
	var fs []string
	for _, e := range es {
		if filepath.Ext(e.Name()) == ".sql" {
			fs = append(fs, e.Name())
		}
	}
	sort.Strings(fs)
	for _, f := range fs {
		b, _ := os.ReadFile(filepath.Join(dir, f))
		_, _ = db.Exec(string(b))
		_, _ = db.Exec(`INSERT INTO schema_migrations(filename) VALUES($1) ON CONFLICT DO NOTHING`, f)
	}
	return nil
}
func main() {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/playground?sslmode=disable"
	}
	db, _ := sql.Open("pgx", dsn)
	defer db.Close()
	_ = apply(db, "migrations")
	fmt.Println("migrations applied")
}
