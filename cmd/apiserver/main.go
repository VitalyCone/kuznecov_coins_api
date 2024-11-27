package main

import (
	"log"
	"os"

	"github.com/VitalyCone/kuznecov_coins_api/internal/app"
	"github.com/VitalyCone/kuznecov_coins_api/internal/app/apiserver"
	"github.com/VitalyCone/kuznecov_coins_api/internal/app/store"
	"gopkg.in/yaml.v3"
)

var (
	configPath  string
	dockerCheck string
	tokenPath   string
)

func init() {
	configPath = "config/apiserver.yaml"
	tokenPath = "config/token.yaml"
	dockerCheck = "DOCKER_ENV"
}

type configData struct {
	ApiAddr     string `yaml:"api_addr"`
	DbUrl       string `yaml:"database_url"`
	DbDockerUrl string `yaml:"database_docker_url"`
}

func main() {
	var configStore *store.Config
	var configServer *apiserver.Config

	cfg := configData{}
	token := app.TokenData{}

	isDocker := os.Getenv(dockerCheck) == "true"

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	dataToken, err := os.ReadFile(tokenPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(dataToken, &token)
	if err != nil {
		log.Fatal(err)
	}

	token.Init()

	if isDocker {
		log.Println("App running in Docker. Using Docker database url")
		configStore = store.NewConfig(cfg.DbDockerUrl)
	} else {
		log.Println("App running without Docker. Using Local database url")
		configStore = store.NewConfig(cfg.DbUrl)
	}

	configServer = apiserver.NewConfig(cfg.ApiAddr)

	store := store.NewStore(configStore)
	server := apiserver.NewAPIServer(configServer, store)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
