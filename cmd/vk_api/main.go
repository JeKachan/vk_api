package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	//"net/http"
	"github.com/JeKachan/vk_api"
)

var accessToken = ""
var userId = ""

func main() {
	fmt.Println("Start programm!")
	fmt.Printf("Insert this link in your browser:\n%s\n", vk_api.FullUrl)
	fmt.Println("Plese enter getted url from page:")

	reader := bufio.NewReader(os.Stdin)
	externalUrl, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	externalUrl = strings.TrimSpace(externalUrl)

	accessToken, err = vk_api.GetAccessTokenFromStr(externalUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	userId, err = vk_api.GetUserId(externalUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//resp, err := http.Get("http://example.com/")
	vk_api.GetAudio(accessToken, userId)

	fmt.Print("End programm!")
	//reader.ReadString('\n')
}