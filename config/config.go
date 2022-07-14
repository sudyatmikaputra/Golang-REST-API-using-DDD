package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var configEnv map[string]string

func init() {
	content, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Println("config.json file not found")
		log.Println(".. using default config")
		configEnv = loadDefaultConfig()
	} else {
		err = json.Unmarshal(content, &configEnv)
		if err != nil {
			log.Println("invalid config.json file")
			log.Println(".. using default config")
			configEnv = loadDefaultConfig()
		}
	}
}

func GetValue(key string) string {
	value, ok := configEnv[key]
	if !ok {
		return ""
	}
	return value
}

func GetEnv(key string) string {
	value, ok := configEnv[key]
	if !ok {
		return ""
	}
	return value
}

func GetLogger() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)

	return log
}
