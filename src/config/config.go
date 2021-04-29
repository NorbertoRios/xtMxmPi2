package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Config struct {
	MainPort           int
	VideoServerPort    int
	MainWANIP          string
	VideoServerWANIP   string
	DBUser             string
	DBPassword         string
	DBName             string
	DBHost             string
	DBPort             string
	DBScheme           string
	DBConnectionString string
	WebServerPort      int
	DBSetMaxIdleConns  int
	DBSetMaxOpenConns  int
}

var config *Config
var m sync.Mutex

func readConfig() (c *Config) {
	dirname, err := os.UserHomeDir()
	file, err := os.OpenFile(dirname+"/config.json", 0, 0644)
	if err != nil {
		log.Fatal("can't open config file: ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	c = &Config{}
	err = decoder.Decode(&c)
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
