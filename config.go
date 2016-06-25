package ghlabels

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/go-github/github"
)

// Config holds the labeler configuration
type Config struct {
	Owner  string          `json:"owner"`
	Repo   string          `json:"repo"`
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
	return &config
}
