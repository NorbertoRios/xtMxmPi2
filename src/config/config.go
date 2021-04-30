package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type Config struct {
	MainPort         int
	VideoServerPort  int
	MainWANIP        string
	VideoServerWANIP string
	WebServerPort    int
	DB               struct {
		DBSetMaxIdleConns  int
		DBSetMaxOpenConns  int
		DBUser             string
		DBPassword         string
		DBName             string
		DBHost             string
		DBPort             string
		DBScheme           string
		DBConnectionString string
	}
}

var config *Config
var m sync.Mutex

func readConfig() (c *Config) {
	confPath := os.Getenv("STREAMAXGOCONFIG")
	if len(confPath) < 1 {
		log.Fatal("no env STREAMAXGOCONFIG read")
	}
	file, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatal("can't read config ", err)
		return nil
	}
	c = &Config{}
	err = json.Unmarshal(file, c)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}
	return c
}

func GetConfig() Config {
	m.Lock()
	if config == nil {
		config = readConfig()
	}
	var c = *config
	m.Unlock()
	return c
}
