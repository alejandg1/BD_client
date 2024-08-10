package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var homedir, err = os.UserHomeDir()
var configdir = filepath.Join(homedir, ".config/bdcli")
var connections = filepath.Join(configdir, "connections.json")
var history = filepath.Join(configdir, "history.json")
var config = filepath.Join(configdir, "config.json")

func Checkdirs() {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if _, err := os.Stat(configdir); os.IsNotExist(err) {
		os.MkdirAll(configdir, 0755)
	}
	if _, err := os.Stat(connections); os.IsNotExist(err) {
		os.Create(connections)
	}
	if _, err := os.Stat(history); os.IsNotExist(err) {
		os.Create(history)
	}
	if _, err := os.Stat(config); os.IsNotExist(err) {
		os.Create(config)
	}
}

func GetConfig() Config {
	var confModel Config
	var file, err = os.ReadFile(config)
	if err != nil {
		fmt.Println("config: ", err)
		os.Exit(1)
	}
	if len(file) > 0 {
		err = json.Unmarshal(file, &confModel)
		if err != nil {
			fmt.Println("config: ", err)
			os.Exit(1)
		}
	}
	return confModel
}

func GetConnections() []Connection {
	var conectModel []Connection
	var file, err = os.ReadFile(connections)
	if err != nil {
		fmt.Println("connections: ", err)
		os.Exit(1)
	}
	if len(file) > 0 {
		err = json.Unmarshal(file, &conectModel)
		if err != nil {
			fmt.Println("connections: ", err)
			os.Exit(1)
		}
	}
	return conectModel
}
func GetHistory() []History {
	var histModel []History
	var file, err = os.ReadFile(history)
	if err != nil {
		fmt.Println("hist: ", err)
		os.Exit(1)
	}
	if len(file) > 0 {
		err = json.Unmarshal(file, &histModel)
		if err != nil {
			fmt.Println("hist: ", err)
			os.Exit(1)
		}

	}
	return histModel
}
