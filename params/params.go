package params

type Params struct {
	Chain string `json:"chain"`
}

type ParamsWithBasicType struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Params  Params `json:"params"`
}

func MockBasicType() ParamsWithBasicType {
	p := ParamsWithBasicType{Name: "mock_name", Address: "0xCdd39B6D1cC4D0a7243b389Ed9356E23Df6240eb", Params: Params{Chain: "ethereum"}}
	return p
}

type BodyName struct {
	Message string `json:"message"`
	Value   string `json:"value"`
}

type Details struct {
	BodyName BodyName `json:"body.name"`
}

type MockRes struct {
	Message string  `json:"message"`
	Details Details `json:"details"`
}

type ReqClassOfYear struct {
	Chain string `json:"chain"`
	Date1 string `json:"date1"`
	Date2 string `json:"date2"`
}

func NewReqClassOfYear(chain string, date1 string, date2 string) ReqClassOfYear {
	return ReqClassOfYear{
		chain,
		date1,
		date2,
	}
}

type ReqBody struct {
	Name            string `json:"name"`
	Address         string `json:"address"`
	IncludeMetadata bool   `json:"includeMetadata"`
	Params          any    `json:"params"`
}

func NewReqBody(name string, address string, params any) ReqBody {
	return ReqBody{
		name,
		address,
		true,
		params,
	}
}

type ResBody struct {
	Name        string   `json:"name"`
	Result      bool     `json:"result"`
	Display     any      `json:"display"`
	Metadata    []string `json:"metadata"`
	RunningCost int      `json:"runningCost"`
}
