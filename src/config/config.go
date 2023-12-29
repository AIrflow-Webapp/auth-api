package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ConfigEnv struct {
	ClientId     string
	ClientSecret string
	Realm        string
}

func LoadEnv() ConfigEnv {
	// Load env
	err := godotenv.Load()
	if err != nil {
		panic("[Error] failed to load env due to: " + err.Error())
	}
	fmt.Println(os.Getenv("CLIENT_ID"))
	config := ConfigEnv{
		ClientId:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Realm:        os.Getenv("KEYCLOAK_REALM"),
	}
	return config
}
