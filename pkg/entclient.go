package pkg

import (
	"context"
	"fmt"
	"language-learning/ent"
	"os"

	_ "github.com/lib/pq"
)

func GetEntClient(ctx context.Context) *ent.Client {
	connectionString := getConnectionString()
	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	if err := client.Schema.Create(ctx); err != nil {
		panic(err)
	}

	return client
}

func getConnectionString() string {
	connectionString := "host=%s port=%s user=%s dbname=%s password=%s sslmode=%s"
	connectionString = fmt.Sprintf(connectionString,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSLMODE"))

	return connectionString
}
