package server

import (
	"fmt"
	"net/http"

	"github.com/locona/ozz/client/infra"
	"github.com/locona/ozz/client/pkg/hydra"
)

func get(url, token string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	bearer := fmt.Sprintf("Bearer %v", token)
	req.Header.Set("Authorization", bearer)
	client := new(http.Client)
	return client.Do(req)
}

func post(url, token string) (*http.Response, error) {
	req, _ := http.NewRequest("POST", url, nil)
	bearer := fmt.Sprintf("Bearer %v", token)
	req.Header.Set("Authorization", bearer)
	client := new(http.Client)
	return client.Do(req)
}

func token() (string, error) {
	cache := hydra.NewCache(infra.Redis)
	return cache.Get()
}
