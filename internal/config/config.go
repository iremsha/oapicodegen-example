package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Log struct {
	Level string `yaml:"level" env:"LOG_LEVEL" env-default:"INFO"`
}

type Sentry struct {
	Dsn string `yaml:"dsn" env:"SENTRY_DSN" env-default:""`
}

type Credentials struct {
	Header string `yaml:"header" env:"PAY_REQUEST_HEADER" env-default:" X-Authorization-Token"`
	APIKey string `yaml:"api_key" env:"PAY_REQUEST_API_KEY" env-default:""`
}

type App struct {
	Env            string      `yaml:"env" env:"APP_ENV" env-default:"local"`
	Port           string      `yaml:"port" env:"APP_PORT" env-default:"9000"`
	MetricsPort    string      `yaml:"metrics_port" env:"APP_METRICS_PORT" env-default:"9100"`
	ReadTimeout    int         `yaml:"read_timeout" env:"APP_READ_TIMEOUT" env-default:"5"`
	ReadBufferSize int         `yaml:"read_buffer_size" env:"APP_READ_BUFFER_SIZE" env-default:"8096"`
	Credentials    Credentials `yaml:"credentials"`
}

type Database struct {
	Host         string `yaml:"host" env:"DB_HOST" env-default:"0.0.0.0"`
	Port         string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	Name         string `yaml:"name" env:"DB_NAME" env-default:"oapicodegen-example"`
	User         string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Pass         string `yaml:"pass" env:"DB_PASS" env-default:"postgres"`
	MaxOpenConns int    `yaml:"maxOpenConns" env:"DB_MAX_OPEN_CONNS" env-default:"50"`
	MaxIdleConns int    `yaml:"maxIdleConns" env:"DB_MAX_IDLE_CONNS" env-default:"10"`
	Schema       string `yaml:"schema" env:"DB_SCHEMA" env-default:"public"`
}

type Config struct {
	Log      Log      `yaml:"log"`
	App      App      `yaml:"app"`
	Sentry   Sentry   `yaml:"sentry"`
	Database Database `yaml:"database"`
}

func Load(path string) (Config, error) {
	cfg := Config{}
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (db *Database) GetDsn() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		db.User,
		db.Pass,
		db.Host,
		db.Port,
		db.Name,
	)
}

func (db *Database) GetMigrateDsn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		db.User,
		db.Pass,
		db.Host,
		db.Port,
		db.Name,
		db.Schema,
	)
}
