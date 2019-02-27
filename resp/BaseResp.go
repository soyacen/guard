package resp

import (
	"encoding/json"
)

type BaseResp struct {
	Error
	Data interface{} `json:"data"`
}

func (resp *BaseResp) ToJson() []byte {
	if resp != nil {
		j, _ := json.Marshal(resp)
		return j
	} else {
		return nil
	}
}
