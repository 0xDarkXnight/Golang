package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps tutorial.")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["GO"] = "Golang"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS shorts for:", languages["JS"])
	fmt.Println("GO shorts for:", languages["GO"])
	fmt.Println("PY shorts for:", languages["PY"])

	delete(languages, "JS")
	fmt.Println("List of all languages:", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
