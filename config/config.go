package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ApiConfig struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
}

func parse(jsonBody []byte, result interface{}) error {
	return json.Unmarshal(jsonBody, &result)
}

//TODO: solve this stupid read from home issue. currently it doesnt get the concept of ~
func Load() ApiConfig {
	data, err := ioutil.ReadFile("/Users/goranojkic/.lfm/config")
	var auth ApiConfig

	if err != nil {
		fmt.Printf("%s", err)
	}

	parse(data, &auth)
	return auth
}
