package common

import "encoding/json"

func ToJsonString(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func FromJsonString(v interface{}, s string) error {
	return json.Unmarshal([]byte(s), v)
}
