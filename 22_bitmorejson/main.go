package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // this won't show
	Tags     []string `json:"tags,omitempty"` // no space after comma is imp
}

func main() {
	fmt.Println("More JSON in Golang")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "learncodeonline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 399, "learncodeonline.in", "cde123", nil},
	}

	// package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
