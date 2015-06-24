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

func check(e error) {
	if e != nil {
		panic(e)
	}
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

func SaveToken(token string) {
	d1 := []byte(token)
	err := ioutil.WriteFile(os.Getenv("HOME")+"/.lfm/token", d1, 0644)
	check(err)
}

func GetToken() string {
	dat, err := ioutil.ReadFile(os.Getenv("HOME") + "/.lfm/token")
	check(err)
	return string(dat)
}

func SaveSession(session string) {
	d1 := []byte(session)
	err := ioutil.WriteFile(os.Getenv("HOME")+"/.lfm/session", d1, 0644)
	check(err)
}

func GetSession() string {
	dat, err := ioutil.ReadFile(os.Getenv("HOME") + "/.lfm/session")
	check(err)
	return string(dat)
}
