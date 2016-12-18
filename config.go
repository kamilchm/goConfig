package goconfig

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const (
	defaultPath       = "./"
	defaultConfigFile = "config.json"
)

// Configuration struct
type Configuration struct {
	Name  string
	Value interface{}
}

// Config instantiate the system settings.
var Config = Configuration{}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Load config file
func Load() (err error) {
	configFile := defaultPath + defaultConfigFile
	file, err := os.Open(configFile)
	if err != nil {
		log.Println("loadConfig open config.json:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		log.Println("loadConfig Decode:", err)
		return
	}
	return
}

// Save config file
func (c *Configuration) Save() (err error) {
	_, err = os.Stat(defaultPath)

	if os.IsNotExist(err) {
		os.Mkdir(defaultPath, 0700)
	}

	configFile := defaultPath + defaultConfigFile

	_, err = os.Stat(configFile)
	if err != nil {
		log.Println(err)
		return
	}

	b, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(defaultConfigFile, b, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
