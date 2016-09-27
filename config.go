package ghlabels

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/go-github/github"
)

const (
	defaultHost = "https://api.github.com/"
)

// Config holds the labeler configuration
type Config struct {
	Host   string          `json:"host"`
	Owner  string          `json:"owner"`
	Repo   string          `json:"repo"`
	Repos  []string        `json:"repos"`
	Labels []*github.Label `json:"labels"`
}

// LoadConfig loads a JSON file from the provided file path
func LoadConfig(fp string) *Config {
	file, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Set the default host if one was not provided
	if config.Host == "" {
		config.Host = defaultHost
	}

	if len(config.Repo) != 0 {
		config.Repos = append(config.Repos, config.Repo)
	}
	return &config
}
