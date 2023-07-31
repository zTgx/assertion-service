package main

import (
)

type Params struct {
	Chain string `json:"chain"`
}

type ParamsWithBasicType struct {
	Name string `json:"name"`
	Address string `json:"address"`
	Params Params `json:"params"`
}

func MockBasicType() ParamsWithBasicType {
	p := ParamsWithBasicType{Name: "mock_name", Address: "0xCdd39B6D1cC4D0a7243b389Ed9356E23Df6240eb", Params: Params{Chain: "ethereum"}}
	return p
}

type BodyName struct {
	Message string `json:"message"`
	Value string `json:"value"`
}

type Details struct {
	BodyName BodyName `json:"body.name"`
}

type MockRes struct {
	Message string `json:"message"`
	Details Details `json:"details"`
}