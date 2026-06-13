package config

import (
	"flag"
	"os"
	"log"
)
import "github.com/ilyakaznacheev/cleanenv"
type HTTPServerConfig struct {
	Addr string 
}

type Config struct{
	Env        string  `yaml:"env" env:"ENV" env-required:"true" env-default:"dev"`
	StoragePath string  `yaml:"storage_path" env-required:"true"`
	HTTPServer HTTPServerConfig `yaml:"http_server"`
}


func MustLoad()*Config{
	var configPath string

	configPath=os.Getenv("CONFIG_PATH")
	if configPath==""{
		flags:=flag.String("config", "config/local.yaml", "path to config file")
		flag.Parse()
		configPath=*flags

		if configPath==""{
			log.Fatal("config path is required")
		}
	}
	if _,err:=os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("config file does not exist at path: %s", configPath)
	}

	var cfg Config 

	err:=cleanenv.ReadConfig(configPath, &cfg)
	if err!=nil{
		log.Fatalf("failed to read config file: %s", err.Error())
	}

	return &cfg

}

