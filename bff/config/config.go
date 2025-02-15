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
	ETCD   ETCD   `yaml:"etcd"`
}

type Server struct {
	User     string `yaml:"user"`
	Favorite string `yaml:"favorite"`
	Feed     string `yaml:"feed"`
	Publish  string `yaml:"publish"`
	Comment  string `yaml:"comment"`
	Relation string `yaml:"relation"`
	Message  string `yaml:"message"`
}

type ETCD struct {
	Addr string `yaml:"addr"`
}

func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	env := getGoEnv()
	prefix := "config"
	filePath := filepath.Join(prefix, filepath.Join(env, "config.yaml"))
	fmt.Println(filePath)
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
