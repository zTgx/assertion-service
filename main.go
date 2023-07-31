package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"log"
)

const ACHAINABLE_NAMES_PATH string = "ACHAINABLE_NAMES.txt"

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

func main() {
	log.SetPrefix("[ACB]: ")
	log.SetFlags(0)

	config := GetServerConfig()
	
	/// Build Body
	basicType := MockBasicType()
	jsonStr, _ := json.Marshal(basicType)

	url := config.Host
	key := config.Key
	fmt.Println("URL: ", url)
	fmt.Println("KEY: ", key)

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", key)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	resBody := &MockRes{}
	derr := json.NewDecoder(res.Body).Decode(resBody)
	if derr != nil {
		panic(derr)
	}

	fmt.Println("[main] res StatusCode: ", res.StatusCode)
	// if res.StatusCode != http.StatusCreated {
	// 	panic(res.Status)
	// }

	details := resBody.Details
	names := details.BodyName.Message
	// fmt.Println("Message:", resBody.Message)
	// fmt.Println("Details:", details)
	content := FilterNames(names)
	WriteToFile(ACHAINABLE_NAMES_PATH, content)
	// ParseMockNames(details.BodyName.Message)
}