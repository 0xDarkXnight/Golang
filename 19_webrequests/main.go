package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://help.websiteos.com/websiteos/example_of_a_simple_html_page.htm"

func main() {
	fmt.Println("WebRequests in Golang")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The Response if of Type: %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
