package config

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"regexp"
)

const (
	configPath = "config/config.json"
)

// Flow center server configuration.
type FlowCenterConfig struct {
	// server port.
	Port int32 `json:"port"`

	// logger config.
	Logger LoggerConfig `json:"logger"`
}

// Log configuration.
type LoggerConfig struct {
	// log file path.
	Path string `json:"path"`
}

// load configuration from json file.
func Load(path string) *FlowCenterConfig {
	var config FlowCenterConfig
	if len(path) == 0 {
		log.Printf("Configuration path is empty in the input, change to use default config path '%s'.", configPath)
		path = configPath
	}

	config_file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Load configuration error, failed to open config file '%s': %s\n", path, err)
	}

	fi, _ := config_file.Stat()

	if fi.Size() == 0 {
		log.Fatalf("Load configuration error, config file (%q) is empty, skipping", path)
	}

	buffer := make([]byte, fi.Size())
	_, err = config_file.Read(buffer)

	buffer, err = stripComments(buffer)
	if err != nil {
		log.Fatalf("Load configuration error, failed to strip comments from json: %s\n", err)
	}

	buffer = []byte(os.ExpandEnv(string(buffer)))

	err = json.Unmarshal(buffer, &config)
	if err != nil {
		log.Fatalf("Load configuration error, failed to strip comments from json: %s\n", err)
	}
	return &config
}

func stripComments(data []byte) ([]byte, error) {
	data = bytes.Replace(data, []byte("\r"), []byte(""), 0)
	lines := bytes.Split(data, []byte("\n"))
	filtered := make([][]byte, 0)

	for _, line := range lines {
		match, err := regexp.Match(`^\s*#`, line)
		if err != nil {
			return nil, err
		}
		if !match {
			filtered = append(filtered, line)
		}
	}

	return bytes.Join(filtered, []byte("\n")), nil
}
