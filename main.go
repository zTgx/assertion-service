package main

import (
	"encoding/json"
	"fmt"
	"golang/acb/client"
	"golang/acb/params"
	"log"
)

const ACHAINABLE_NAMES_PATH string = "./config/ACHAINABLE_NAMES.txt"

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

func main() {
	log.SetPrefix("[ACB]: ")
	log.SetFlags(0)

	client := client.New()
	res := client.Send()
	fmt.Println("res.body: ", res.Body)
	resBody := &params.MockRes{}

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

func MockBasicType() {
	panic("unimplemented")
}

func GetServerConfig() {
	panic("unimplemented")
}
