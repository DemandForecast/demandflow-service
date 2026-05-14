package utils

import (
	"DemandFlow-Service/dto"
	"os"
)

func LoadAuthConfig() dto.AuthConfig {
	return dto.AuthConfig{
		AUTH0_DOMAIN:   os.Getenv("AUTH0_DOMAIN"),
		AUTH0_AUDIENCE: os.Getenv("AUTH0_AUDIENCE"),
	}
}
