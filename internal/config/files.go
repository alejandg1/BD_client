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
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetData[data []Connection | Config | []History]() (data, error) {
	var path string
	var resp data
	var err error = nil
	var file []byte
	switch any(resp).(type) {
	case []Connection:
		path = connections
	case []History:
		path = history
	case Config:
		path = config
	default:
		err = fmt.Errorf("Invalid type: %T", resp)
	}
	file, err = os.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading file: %s", path)
	}
	if len(file) > 0 {
		err = json.Unmarshal(file, &resp)
	}
	return resp, err
}
