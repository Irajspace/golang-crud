package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServerConfig struct {
	Addr string `yaml:"address" env:"http_server" env-required:"true" env-default:":8080"`
}

type Config struct {
	Env         string           `yaml:"env" env:"ENV" env-required:"true" env-default:"dev"`
	StoragePath string           `yaml:"storage_path" env-required:"true"`
	HTTPServer  HTTPServerConfig `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "config/local.yaml", "path to config file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("config path is required")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist at path: %s", configPath)
	}

	var cfg Config
	log.Printf("loading config from file: %s\n", configPath)
	err := cleanenv.ReadConfig(configPath, &cfg)
	log.Printf("config loaded successfully: %+v\n", cfg)
	if err != nil {
		log.Fatalf("failed to read config file: %s", err.Error())
	}

	return &cfg

}
