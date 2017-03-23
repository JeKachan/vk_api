package vkapi

import "fmt"

//vk application id
const ClientID = "5637412"
const FormatURL = "https://oauth.vk.com/authorize?client_id=%s&display=page&redirect_uri=https://oauth.vk.com/blank.html&scope=audio&response_type=token&v=5.52"
const QueryVkAPITemplate = "https://api.vk.com/method/{{.methodName}}?{{.parameters}}&access_token={{.accessToken}}&v={{.version}}"

var FullURL = fmt.Sprintf(FormatURL, ClientID)
