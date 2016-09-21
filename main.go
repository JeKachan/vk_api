package main

import (
	"fmt"
	//"github.com/jekachan/vk_api/audio"
	"bufio"
	"os"
	"strings"
)

//vk application id
const client_id  = "5637412"
const format_url = "https://oauth.vk.com/authorize?client_id=%s&display=page&redirect_uri=https://oauth.vk.com/blank.html&scope=audio&response_type=token&v=5.52"

var full_url = fmt.Sprintf(format_url, client_id)

func main() {
	fmt.Println("Start programm!")

	fmt.Printf("Insert this link in your browser:\n%s\n", full_url)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Plese enter getted url from page:")

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	text = strings.TrimSpace(text)

	fmt.Println("End programm!")
	reader.ReadString('\n')
}
