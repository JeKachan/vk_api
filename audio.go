package vkapi

import (
	"bytes"
	"fmt"
	"html/template"
)

func GetAudio(accessToken string, userID string) (string, error) {
	result := ""
	vap := NewParams()
	vap.accessToken = accessToken
	vap.methodName = "audio.get"
	vap.parameters = "owner_id=" + userID

	tmpl, err := template.New("VkApiUrl").Parse(QueryVkAPITemplate)
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}

	var buf bytes.Buffer
	tmpl.Execute(&buf, vap)
	result = buf.String()

	return result, nil
}
