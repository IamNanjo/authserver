package config

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// Struct fields env and/or flag tags get parsed. Flag takes precedence over env.
type Config struct {
	Address     *string `required:"true" flag:"addr,Listen address (required)" env:"AUTHSERVER_ADDRESS"`
	DatabaseURI *string `flag:"db,Database URI (default <path of executable>/authserver.db)" env:"AUTHSERVER_DB"`
}

type fields struct {
	flags    map[string]*field
	env      map[string]*field
	required []*field
}
type field struct {
	t reflect.StructField
	v reflect.Value
}

// Populate Config with CLI flags, .env file and environment variables in that priority order.
// Will exit program with code 1 if required fields have no value missing
func Parse() *Config {
	c := &Config{}
	f := c.parseFields()

	return c.parseFlags(f).parseDotEnv(f).parseEnv(f).checkRequired(f)
}

func (c *Config) parseFields() *fields {
	fields := &fields{
		flags:    map[string]*field{},
		env:      map[string]*field{},
		required: []*field{},
	}

	config := reflect.ValueOf(c).Elem()
	configType := config.Type()

	for i := range configType.NumField() {
		fieldType := configType.Field(i)
		fieldValue := config.Field(i)

		field := &field{t: fieldType, v: fieldValue}

		flagTag, flagTagSet := fieldType.Tag.Lookup("flag")
		if flagTagSet {
			fields.flags[flagTag] = field
		}

		envTag, envTagSet := fieldType.Tag.Lookup("env")
		if envTagSet {
			fields.env[envTag] = field
		}

		_, required := fieldType.Tag.Lookup("required")
		if required {
			fields.required = append(fields.required, field)
		}
	}

	return fields
}

// Will exit program with code 1 if required fields are not provided
func (c *Config) checkRequired(fields *fields) *Config {
	for _, field := range fields.required {
		if field.v.IsNil() || field.v.Elem().String() == "" {
			fmt.Fprintf(os.Stderr, "Required option %s not set\n", field.t.Name)
			os.Exit(1)
		}
	}
	return c
}

func (c *Config) parseFlags(fields *fields) *Config {
	for _, field := range fields.flags {
		flagTag, isSet := field.t.Tag.Lookup("flag")
		if !isSet {
			continue
		}

		flagProperties := strings.Split(flagTag, ",")
		description := ""

		if len(flagProperties) > 1 {
			description = flagProperties[1]
		}

		field.v.Set(reflect.ValueOf(flag.String(flagProperties[0], "", description)))
	}

	flag.Parse()
	return c
}

func (c *Config) parseDotEnv(fields *fields) *Config {
	file, err := os.Open(".env")
	if err != nil {
		return c
	}
	defer file.Close()

	// Read line by line
	scanner := bufio.NewScanner(file)

	envVars := map[string]string{}

	// Parse and collect environment variable keys and values
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || line[0] == '#' {
			continue
		}

		// Cut everything after comment (#).
		// Ignores # if it is not the first character or if there is no space around it.
		for i, char := range line {
			if char == '#' && (i == 0 || line[i-1] == ' ') {
				line = line[0 : i-1]
			}
		}

		arg := strings.Split(line, "=")
		envVars[arg[0]] = arg[1]
	}

	// Check if Config fields use any of the parsed variables
	for key, field := range fields.env {
		// Already set
		if !field.v.IsNil() && field.v.Elem().String() != "" {
			continue
		}

		val := envVars[key]
		if val == "" {
			continue
		}

		field.v.Set(reflect.ValueOf(&val))
	}

	return c
}

func (c *Config) parseEnv(fields *fields) *Config {
	for key, field := range fields.env {
		// Already set
		if !field.v.IsNil() && field.v.Elem().String() != "" {
			continue
		}

		val, exists := os.LookupEnv(key)
		if !exists {
			continue
		}

		field.v.Set(reflect.ValueOf(&val))
	}

	return c
}
