package vkapi

import (
	"strings"
	"errors"
)

// Get Access Token String from source string
func GetAccessTokenFromStr(s string) (string, error) {
	return getParamByName(s, "access_token")
}

/**
 * Get User Id from url
 */
func GetUserId(s string) (string, error) {
	return getParamByName(s, "user_id")
}

func getParamByName(s string, name string) (string, error) {
	queryString := strings.Split(s, "#")[1]
	param := ""
	if queryString == "" {
		return param, errors.New("Error: query string is empty!")
	}
	queryStringItems := strings.Split(queryString, "&")
	for i := 0; i < len(queryStringItems); i++ {
		if strings.Index(queryStringItems[i], name) == -1 {
			continue
		}
		param = strings.Split(queryStringItems[i], "=")[1]
	}
	return param, nil
}

