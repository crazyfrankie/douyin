package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env    string
	Server Server `yaml:"server"`
	MySQL  MySQL  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	Kafka  Kafka  `yaml:"kafka"`
	Etcd   ETCD   `yaml:"etcd"`
	RPC    RPC    `yaml:"grpc"`
	JWT    JWT    `yaml:"jwt"`
}

type Server struct {
	Address string `yaml:"address"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

type Redis struct {
	Address string `yaml:"address"`
}

type Kafka struct {
	Address string `yaml:"address"`
}

type RPC struct {
	Address string `yaml:"address"`
}
type ETCD struct {
	Address string `yaml:"address"`
}

type JWT struct {
	SecretKey string `yaml:"secretKey"`
}

func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	env := getGoEnv()
	prefix := "config"
	filePath := filepath.Join(prefix, filepath.Join(env, "config.yaml"))
	viper.SetConfigFile(filePath)

	if err := viper.ReadInConfig(); err != nil {
		panic(errors.New("failed read config"))
	}

	conf = new(Config)
	err := viper.Unmarshal(conf)
	if err != nil {
		panic(errors.New("failed parse config file"))
	}

	conf.Env = env
	fmt.Printf("%v\n", conf)
}

func getGoEnv() string {
	env, ok := os.LookupEnv("GO_ENV")
	if !ok {
		return "test"
	}

	return env
}
