package pkg

import (
	"fmt"
	"language-learning/ent"
	"os"

	"github.com/spf13/viper"
)

func GetEntClient() *ent.Client {
	connectionString := getConnectionString()
	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	return client
}

func getConnectionString() string {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = viper.GetString("db.host")
	}

	connectionString := "host=%s port=%d user=%s dbname=%s password=%s sslmode=%s"
	connectionString = fmt.Sprintf(connectionString,
		dbHost,
		viper.GetInt("db.port"),
		viper.GetString("db.user"),
		viper.GetString("db.name"),
		viper.GetString("db.password"),
		viper.GetString("db.sslmode"))

	return connectionString
}
