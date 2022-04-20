package utils

import (
	"io/ioutil"
	"strings"

	"github.com/go-resty/resty/v2"
)

func ContentOfPath(path string) (data []byte, err error) {

	isRemote := strings.HasPrefix(path, "http")
	if isRemote {
		var res *resty.Response
		res, err = resty.New().NewRequest().Get(path)
		if err != nil {
			return
		}
		data = res.Body()
		_ = res.RawBody().Close()
		return

	} else {
		data, err = ioutil.ReadFile(path)
	}
	return
}
