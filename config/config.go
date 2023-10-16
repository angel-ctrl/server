package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

// Configs is a struct that holds the configuration values
type Env struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`

	User     string `mapstructure:"user"`
	Pass     string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	HostDB   string `mapstructure:"hostDB"`
	PortDB   string `mapstructure:"portDB"`

	SecretKeyJWT string `mapstructure:"SECRET_KEY_JWT"`

	PongWait       int `mapstructure:"PongWait"`
	MaxMessageSize int `mapstructure:"MaxMessageSize"`

	Secret string `mapstructure:"Secret"`
}

func CreateConfigs() *Env {

	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		//log.Fatal("Can't find the file .env : ", err)

		env.Host = os.Getenv("HOST")
		env.Port = os.Getenv("PORT")

		env.User = os.Getenv("user")
		env.Pass = os.Getenv("password")
		env.Database = os.Getenv("database")
		env.HostDB = os.Getenv("hostDB")
		env.PortDB = os.Getenv("portDB")

		env.SecretKeyJWT = os.Getenv("SECRET_KEY_JWT")

		env.PongWait, err = strconv.Atoi(os.Getenv("PongWait"))
		if err != nil{
			log.Fatal("Can't use PongWait: ", err)
		}
		
		env.MaxMessageSize, err = strconv.Atoi(os.Getenv("MaxMessageSize"))
		if err != nil{
			log.Fatal("Can't use PongWait: ", err)
		}

		env.Secret = os.Getenv("Secret")

		return &env
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}
