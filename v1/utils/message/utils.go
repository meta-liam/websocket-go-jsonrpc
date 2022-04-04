package message

import (
	"encoding/json"
	"fmt"
)

type MsgStruct struct {
	Version string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  string `json:"params"`
	ID      int64  `json:"id"`
	Service string `json:"service"`
}

func ToJson(b []byte) {
	//json str 转map
	var dat map[string]interface{}
	if err := json.Unmarshal(b, &dat); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(dat)
		fmt.Println(dat["service"])
	}
}

func ToMap(b []byte) map[string]interface{} {
	//json str 转map
	var data map[string]interface{}
	if err := json.Unmarshal(b, &data); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(data)
		fmt.Println(data["service"])
	}
	return data
}
