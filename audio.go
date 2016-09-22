package vk_api

import (
	"html/template"
	"fmt"
	"bytes"
)

func GetAudio(accessToken string, userId string) {
	vap := NewVkApiParams()
	vap.ACCESS_TOKEN = accessToken
	vap.METHOD_NAME = "audio.get"
	vap.PARAMETERS = "owner_id=" + userId
	tmpl, err := template.New("VkApiUrl").Parse(QueryVkApiTemplate)
	if err != nil {
		fmt.Println(err.Error())
	}

	var buf bytes.Buffer
	tmpl.Execute(&buf, vap)
	s := buf.String()
	fmt.Println(s)

}