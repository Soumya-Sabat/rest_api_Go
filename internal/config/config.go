package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

// env-default:"production"
type Config struct {
	Env         string               `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string               `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"` //embedding another struct
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == ""{
		//to see if someone passed the flag in the runtime args
		flags := flag.String("config","","path to the configuration file")
		flag.Parse()
		configPath = *flags
		if configPath==""{
			log.Fatal("Config path is not set")
		}
	}

	//chk wheather that the file is availiable or not 
	if _,err := os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("config file does not exist: %s",configPath)
	}

	var cfg Config
	err:=cleanenv.ReadConfig(configPath,&cfg)
	if err!=nil{
		log.Fatalf("cannot read the config file : %s",err.Error())
	}
	return &cfg 
}