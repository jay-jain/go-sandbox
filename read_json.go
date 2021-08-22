package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// type Message struct {
// 	Name string
// 	Body string
// 	Time int
// }
const data = "vulns.json"

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open(data)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened", data)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result)
}
