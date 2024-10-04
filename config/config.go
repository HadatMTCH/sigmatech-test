package config

import (
	"log"
	"os"

	"base-api/constants"

	"github.com/spf13/viper"
)

type Config struct {
	Server    ServerConfig
	DB        DBConfig
	Redis     RedisServer
	JWTConfig JWTConfig
	S3        S3Configuration
	Mailer    Mailer
	FCM       FCM
}

type ServerConfig struct {
	Addr            string
	WebsocketAddr   string
	WriteTimeout    int
	ReadTimeout     int
	GraceFulTimeout int
	Registration    bool
}

type DBConfig struct {
	Name            string
	Host            string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime int
}

type RedisServer struct {
	Addr     string
	Password string
	Timeout  int
	MaxIdle  int
}

type JWTConfig struct {
	Issuer            string
	Secret            string
	TokenLifeTimeHour int
}

type S3Configuration struct {
	Key        string
	Secret     string
	Region     string
	Bucket     string
	RootFolder string
	PublicUrl  string
}

type Mailer struct {
	Server     string
	Port       int
	Username   string
	Password   string
	UseTls     bool
	Sender     string
	MaxAttempt int
}

type FCM struct {
	ProjectID  string
	KeyFileDir string
}

func InitConfig() Config {
	viper.SetConfigName(".env")
	if os.Getenv("ENV") == constants.ENV_STAGING {
		viper.SetConfigName(".env-" + constants.ENV_STAGING)
	}

	if os.Getenv("ENV") == constants.ENV_PRODUCTION {
		viper.SetConfigName(".env-" + constants.ENV_PRODUCTION)
	}

	viper.AddConfigPath(".")

	var configuration Config

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return configuration
}
