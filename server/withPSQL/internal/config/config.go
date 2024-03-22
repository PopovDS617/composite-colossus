package config

type PGConfig interface {
	DSN() string
}

type HTTPConfig interface {
	Port() string
}
