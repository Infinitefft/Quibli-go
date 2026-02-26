package main

import (
	"context"
	"log"

	"backend/internal/ent/migrate"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

func main() {
	dsn := "host=localhost port=5433 user=postgres password=123456 dbname=postgres sslmode=disable"

	db, err := sql.Open(dialect.Postgres, dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	_, err = db.ExecContext(ctx, "CREATE EXTENSION IF NOT EXISTS vector")
	if err != nil {
		log.Printf("warning: could not create vector extension: %v", err)
	}

	schema := migrate.NewSchema(db)

	if err := schema.Create(ctx); err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	log.Println("Migration completed successfully!")
}
