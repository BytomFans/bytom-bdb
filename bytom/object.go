package bytom

import (
	"encoding/json"
	"fmt"
)

type JSONRPCObject struct {
	Params interface{} `json:"params"`
}

func (jrpc *JSONRPCObject) AsJsonString() string {
	resultBytes, err := json.Marshal(jrpc)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(resultBytes)
}
