package common

import "encoding/json"

type CommonResponse struct {
	*BaseResponse
	*actionRequest
}

func NewCommonResponse() (response *CommonResponse) {
	response = &CommonResponse{
		BaseResponse:  &BaseResponse{},
		actionRequest: &actionRequest{},
	}
	return
}

func (r *CommonResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, r.actionRequest)
}

func (r *CommonResponse) GetBody() *actionRequest {
	return r.actionRequest
}
