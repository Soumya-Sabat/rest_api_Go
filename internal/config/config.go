package config

import (
	"flag"
	"os"
)

type HTTPServer struct {
	Addr string
}

// env-default:"production"
type Config struct {
	Env         string               `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string               `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"` //embedding another struct
}

func MustLoad() {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == ""{
		//to see if someone passed the flag in the runtime args
		flags := flag.String("config","","path to the configuration file")
		flag.Parse()
		configPath = *flags
	}
}