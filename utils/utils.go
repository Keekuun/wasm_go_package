package utils

import "encoding/json"

type Result struct {
	OK   bool
	Err  string
	Data any
}

func PackResult(ok bool, err error, data any) string {
	_err := ""
	if err != nil {
		_err = err.Error()
	}

	res := Result{
		OK:   ok,
		Err:  _err,
		Data: data,
	}

	// json 序列化
	resByte, _ := json.Marshal(res)

	return string(resByte)
}
