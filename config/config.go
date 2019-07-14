package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Credenciais e servidor do banco
type Config struct {
	Servidor string
	Database string
}

// Lê e faz o parse do arquivo de configuração
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
