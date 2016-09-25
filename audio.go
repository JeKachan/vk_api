package vkapi

import (
	"html/template"
	"fmt"
	"bytes"
)

func GetAudio(accessToken string, userId string) (string, error) {
	result := ""
	vap := NewVkApiParams()
	vap.ACCESS_TOKEN = accessToken
	vap.METHOD_NAME = "audio.get"
	vap.PARAMETERS = "owner_id=" + userId

	tmpl, err := template.New("VkApiUrl").Parse(QueryVkApiTemplate)
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}

	var buf bytes.Buffer
	tmpl.Execute(&buf, vap)
	result = buf.String()

	return result, nil
}