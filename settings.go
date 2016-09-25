package vkapi

import "fmt"

//vk application id
const Client_id  = "5637412"
const Format_url = "https://oauth.vk.com/authorize?client_id=%s&display=page&redirect_uri=https://oauth.vk.com/blank.html&scope=audio&response_type=token&v=5.52"
const QueryVkApiTemplate = "https://api.vk.com/method/{{.METHOD_NAME}}?{{.PARAMETERS}}&access_token={{.ACCESS_TOKEN}}&v={{.V}}"

var FullUrl = fmt.Sprintf(Format_url, Client_id)


