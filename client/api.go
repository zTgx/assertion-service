package client

import (
	"encoding/json"
	"golang/acb/config"
	"golang/acb/params"
	"golang/acb/utils"
)

func updateNames(resBody *params.MockRes) {
	details := resBody.Details
	names := details.BodyName.Message
	// fmt.Println("Message:", resBody.Message)
	// fmt.Println("Details:", details)
	content := utils.FilterNames(names)
	utils.WriteToFile(config.GetServerConfig().NamePath, content)
	// ParseMockNames(details.BodyName.Message)
}

func UpdateMockName() {
	client := New()
	res := client.Send()

	resBody := &params.MockRes{}
	derr := json.NewDecoder(res.Body).Decode(resBody)
	if derr != nil {
		panic(derr)
	}

	// fmt.Println("[main] res StatusCode: ", res.StatusCode)
	// if res.StatusCode != http.StatusCreated {
	// 	panic(res.Status)
	// }

	updateNames(resBody)
}
