package config

import (
	"github.com/IamNanjo/go-flagenv"
	"github.com/IamNanjo/go-logging/pkg/format"
)

var Parsed *config

type configWrapper struct {
	Config *config `env:"AUTHSERVER_"`
}

// Struct fields env and/or flag tags get parsed. Flag takes precedence over env.
type config struct {
	Address      string         `flag:"addr"     env:"ADDRESS"   required:"true"              desc:"Listen address"`
	DatabasePath string         `flag:"db"       env:"DB"        default:"/tmp/authserver/authserver.db" desc:"Database path"`
	WebAuthn     WebAuthnConfig `flag:"webAuthn" env:"WEBAUTHN_"`
}

type WebAuthnConfig struct {
	Id          string   `flag:"Id"          env:"ID"           desc:"Base domain for the service"`
	DisplayName string   `flag:"DisplayName" env:"DISPLAY_NAME" desc:"Human readable name for the service" default:"Authentication Service"`
	Origins     []string `flag:"Origins"     env:"ORIGINS"      desc:"Full origin URL to the service"`
}

func Parse() error {
	parsed := new(configWrapper)
	parsed.Config = new(config)
	if err := flagenv.Parse(parsed); err != nil {
		return format.Err("Flagenv parsing failed %w", err)
	}
	Parsed = parsed.Config
	return nil
}
