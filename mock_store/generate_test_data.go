package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/asstrahanec/go-clickhouse/pkg/repository"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func randomEventType() string {
	eventTypes := []string{"login", "purchase", "logout", "update_profile", "delete_account"}
	return eventTypes[rand.Intn(len(eventTypes))]
}

func randomUserID() int64 {
	return rand.Int63n(10000) // Генерация случайного userID
}

func randomEventTime() time.Time {
	min := time.Now().AddDate(-1, 0, 0) // Минимальная дата - год назад
	max := time.Now()                   // Максимальная дата - текущее время
	delta := max.Sub(min)
	sec := rand.Int63n(int64(delta.Seconds()))
	return min.Add(time.Duration(sec) * time.Second)
}

func randomPayload() string {
	payloads := []string{
		`{"details": "user logged in"}`,
		`{"details": "user made a purchase"}`,
		`{"details": "user logged out"}`,
		`{"details": "user updated profile"}`,
		`{"details": "user deleted account"}`,
	}
	return payloads[rand.Intn(len(payloads))]
}

func main() {
	count := flag.Int("count", 10, "Number of random events to generate")
	flag.Parse()

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %s", err.Error())
	}

	db, err := repository.NewClickHouseDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		log.Fatalf("error initializing db: %s", err.Error())
	}

	insertEventQuery := `INSERT INTO events (eventType, userID, eventTime, payload)
	                     VALUES (?, ?, ?, ?)`

	for i := 0; i < *count; i++ {
		eventType := randomEventType()
		userID := randomUserID()
		eventTime := randomEventTime()
		payload := randomPayload()

		_, err := db.Exec(insertEventQuery, eventType, userID, eventTime, payload)
		if err != nil {
			log.Fatalf("error inserting random test data: %s", err.Error())
		}
	}

	fmt.Printf("Successfully inserted %d random events\n", *count)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
