package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://darkxnight.bat:3000/glide?name=bruce&surname=wayne"

func main() {
	fmt.Println("UrlRequests in Golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams) // url.Values, it is basically key value pair

	fmt.Println(qparams["name"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "darkxnight.bat",
		Path:     "/glide",
		RawQuery: "fullname=brucewayne",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
