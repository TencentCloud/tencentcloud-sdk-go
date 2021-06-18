package common

import "encoding/json"

type CommonResponse struct {
	*BaseResponse
	*actionParameters
}

func NewCommonResponse() (response *CommonResponse) {
	response = &CommonResponse{
		BaseResponse:     &BaseResponse{},
		actionParameters: &actionParameters{},
	}
	return
}

func (r *CommonResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, r.actionParameters)
}

func (r *CommonResponse) GetBody() *actionParameters {
	return r.actionParameters
}
