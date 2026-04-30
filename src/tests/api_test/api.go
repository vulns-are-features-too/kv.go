// Package api_test for API tests
package api_test

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

const serverURL = "http://127.0.0.1:43219"

type testAPI struct {
	server string
}

func makeAPI() testAPI {
	return testAPI{
		server: serverURL,
	}
}

func (api testAPI) req(endpoint, data string) (string, error) {
	resp, err := http.Post(
		fmt.Sprintf("%s/%s", api.server, endpoint),
		"application/x-www-form-urlencoded",
		strings.NewReader(data),
	)
	if err != nil {
		return "", fmt.Errorf("API error: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading body: %w", err)
	}

	return string(body), nil
}

func (api testAPI) get(key string) (string, error) {
	return api.req("get", key)
}

func (api testAPI) getKeys() ([]string, error) {
	res, err := api.req("getkeys", "")
	if err != nil {
		return make([]string, 0), err
	}

	return strings.Split(res, "\n"), nil
}

func (api testAPI) set(key, value string) error {
	_, err := api.req("set", fmt.Sprintf("%s=%s", key, value))

	return err
}

func randKey(chars int) string {
	const (
		keyAlphabet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		keyAlphabetLen = len(keyAlphabet)
	)

	sb := strings.Builder{}
	sb.Grow(chars)

	for range chars {
		//nolint:gosec
		c := keyAlphabet[rand.Int63()%int64(keyAlphabetLen)]
		sb.WriteByte(c)
	}

	return sb.String()
}
