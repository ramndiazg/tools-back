package main

import (
	"context"
	"log"
	"tools-back/ent"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=admin dbname=tools-back password=1234 sslmode=disable"

	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Println("Migration completed successfully!")
}