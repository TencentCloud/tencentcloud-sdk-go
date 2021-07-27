package ini

import (
	"fmt"
	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"strconv"
)

type value struct {
	raw string
}

func (v *value) Int() (int, error) {
	i, e := strconv.Atoi(v.raw)
	if e != nil {
		msg := fmt.Sprintf("failed to parsing %s to Int: %s", v.raw, e.Error())
		e = tcerr.NewTencentCloudSDKError(iniErr, msg, "")
	}
	return i, e
}

func (v *value) Int64() (int64, error) {
	i, e := strconv.ParseInt(v.raw, 10, 64)
	if e != nil {
		msg := fmt.Sprintf("failed to parsing %s to Int64: %s", v.raw, e.Error())
		e = tcerr.NewTencentCloudSDKError(iniErr, msg, "")
	}
	return i, e
}

func (v *value) String() string {
	return v.raw
}

func (v *value) Bool() (bool, error) {
	switch v.raw {
	case "1", "t", "T", "true", "TRUE", "True", "YES", "yes", "Yes", "y", "ON", "on", "On":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False", "NO", "no", "No", "n", "OFF", "off", "Off":
		return false, nil
	}
	errorMsg := fmt.Sprintf("failed to parsing \"%s\" to Bool: invalid syntax", v.raw)
	return false, tcerr.NewTencentCloudSDKError(iniErr, errorMsg, "")
}

func (v *value) Float32() (float32, error) {
	f, e := strconv.ParseFloat(v.raw, 32)
	if e != nil {
		msg := fmt.Sprintf("failed to Parse %s to Float32: %s", v.raw, e.Error())
		e = tcerr.NewTencentCloudSDKError(iniErr, msg, "")
	}
	return float32(f), e
}
func (v *value) Float64() (float64, error) {
	f, e := strconv.ParseFloat(v.raw, 64)
	if e != nil {
		msg := fmt.Sprintf("failed to Parse %s to Float64: %s", v.raw, e.Error())
		e = tcerr.NewTencentCloudSDKError(iniErr, msg, "")
	}

	return f, e
}
