package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang/acb/config"
	"golang/acb/params"
	"net/http"
)

type AClient struct {
	Config config.Config `json:"config"`
}

func New() AClient {
	config := config.GetServerConfig()
	fmt.Println("Config: ", config)

	// url := config.Host
	// key := config.Key
	// fmt.Println("URL: ", url)
	// fmt.Println("KEY: ", key)

	return AClient{config}
}

func (aclient *AClient) Send() *http.Response {
	url := aclient.Config.Host
	key := aclient.Config.Key

	body := Body()
	r, err := http.NewRequest("POST", url, body)
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

	// defer res.Body.Close()
	return res
}

func Body() *bytes.Buffer {
	basicType := params.MockBasicType()
	jsonStr, _ := json.Marshal(basicType)
	return bytes.NewBuffer(jsonStr)
}
