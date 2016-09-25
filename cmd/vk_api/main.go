package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	//"net/http"
	"github.com/JeKachan/vkapi"
)

var accessToken = ""
var userId = ""

func main() {
	fmt.Println("Start programm!")
	fmt.Printf("Insert this link in your browser:\n%s\n", vkapi.FullUrl)
	fmt.Println("Plese enter getted url from page:")

	reader := bufio.NewReader(os.Stdin)
	externalUrl, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	externalUrl = strings.TrimSpace(externalUrl)

	accessToken, err = vkapi.GetAccessTokenFromStr(externalUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	userId, err = vkapi.GetUserId(externalUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//resp, err := http.Get("http://example.com/")
	vkapi.GetAudio(accessToken, userId)

	fmt.Print("End programm!")
	//reader.ReadString('\n')
}