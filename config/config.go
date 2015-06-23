package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ApiConfig struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
}

func parse(jsonBody []byte, result interface{}) error {
	return json.Unmarshal(jsonBody, &result)
}

func Load() ApiConfig {
	data, err := ioutil.ReadFile(os.Getenv("HOME") + "/.lfm/config")
	var auth ApiConfig

	if err != nil {
		fmt.Printf("%s", err)
	}

	parse(data, &auth)
	return auth
}
