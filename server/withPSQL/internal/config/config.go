package config

type PGConfig interface {
	DSN() string
}
