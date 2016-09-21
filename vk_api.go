package vk_api

import (
	"strings"
	"errors"
)

// Get Access Token String from source string
func GetAccessTokenFromStr(s string) (string, error) {
	queryString := strings.Split(s, "#")[1]
	accessToken := ""
	if queryString == "" {
		return accessToken, errors.New("Error: query string is empty!")
	}
	queryStringItems := strings.Split(queryString, "&")
	for i := 0; i < len(queryStringItems); i++ {
		if strings.Index(queryStringItems[i], "access_token") == -1 {
			continue
		}
		accessToken = strings.Split(queryStringItems[i], "=")[1]
	}
	return accessToken, nil
}

