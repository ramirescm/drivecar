package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Auth(path, user, password string) error {
	creds := credentials{user, password}

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(creds)
	if err != nil {
		return err
	}

	response, err := doRequest("POST", path, &body, false)
	if err != nil {
		return err
	}

	return createToken(response.Body)
}

type cacheToken struct {
	Token string `json:"token"`
}

func createToken(body io.Reader) error {
	token, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	file, err := os.Create(".cacheToken")
	if err != nil {
		return err
	}

	cache := cacheToken{string(token)}
	data, err := json.Marshal(&cache)
	if err != nil {
		return err
	}

	_, err = file.Write(data)

	return err
}

func readCacheToken() (string, error) {
	data, err := os.ReadFile(".cacheToken")

	var cache cacheToken
	err = json.Unmarshal(data, &cache)
	if err != nil {
		return "", err
	}

	return cache.Token, nil
}
