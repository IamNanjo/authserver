package config

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

var addrFlag = "addr"
var dbFlag = "db"
var addrEnvKey = "AUTHSERVER_ADDRESS"
var dbEnvKey = "AUTHSERVER_DB"

// Struct fields env and/or flag tags get parsed
type Config struct {
	Address     string
	DatabaseURI string
}

// Populate AppConfig with CLI flags, .env file and environment variables in that priority order.
func (c *Config) ParseConfig() *Config {
	return c.parseFlags().parseDotEnv().parseEnv()
}

func (c *Config) parseFlags() *Config {
	flag.StringVar(&c.Address, addrFlag, "", "Listen address")
	flag.StringVar(&c.DatabaseURI, dbFlag, "", "Database path")
	flag.Parse()

	return c
}

func (c *Config) parseDotEnv() *Config {
	file, err := os.Open(".env")
	if err != nil {
		return c
	}
	defer file.Close()

	// Scan
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Cut everything after comment (#).
		// Ignores # if it is not the first character or if there is no space around it.
		for i, char := range line {
			if char == '#' && (i == 0 || line[i-1] == ' ') {
				line = line[0:i]
			}
		}

		arg := strings.Split(line, "=")

		// Skip malformed lines
		if len(arg) != 2 {
			continue
		}

		// Set valid variables
		switch arg[0] {
		case addrEnvKey:
			if c.Address == "" {
				c.Address = arg[1]
			}
		case dbEnvKey:
			if c.DatabaseURI == "" {
				c.DatabaseURI = arg[1]
			}
		}
	}

	return c
}

func (c *Config) parseEnv() *Config {
	if c.Address == "" {
		c.Address = os.Getenv(addrEnvKey)
	}
	if c.DatabaseURI == "" {
		c.DatabaseURI = os.Getenv(dbEnvKey)
	}

	return c
}
