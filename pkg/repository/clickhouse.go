package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewClickHouseDB(cfg Config) (*sqlx.DB, error) {
	// Формируем строку подключения без параметра database
	dsn := fmt.Sprintf("tcp://%s:%s?username=%s&password=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password)

	// Подключаемся к базе данных
	db, err := sqlx.Open("clickhouse", dsn)
	if err != nil {
		return nil, err
	}

	// Проверяем соединение
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Выполняем команду для выбора базы данных
	_, err = db.Exec(fmt.Sprintf("USE %s", cfg.DBName))
	if err != nil {
		return nil, fmt.Errorf("failed to set database: %w", err)
	}

	return db, nil
}
