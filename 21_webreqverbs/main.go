package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("WebRequestVerbs in Golang")
	PerformPostRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content Length is:", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is:", byteCount)
	fmt.Println(responseString.String())

	fmt.Println(content)
	fmt.Println(string(content))
}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	// fake JSON payload

	requestBody := strings.NewReader(`
		{
			"dialogue":"I'm Vengeance",
			"batrangs":"10",
			"platform":"darkxnight.bat"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
