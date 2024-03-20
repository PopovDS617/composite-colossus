package config

import (
	"errors"
	"fmt"
	"os"
)

const (
	dsnEnvName = "PG_DSN"
)

type pgConfig struct {
	dsn string
}

func NewPGConfig() (PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	fmt.Println(dsn)
	if len(dsn) == 0 {
		return nil, errors.New("pg dsn not found")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
