package utils

import "testing"

func TestData(t *testing.T) {

	data, err := ContentOfPath("https://imovie.seungyu.cn/api/v2/yi/master/")
	t.Log(data)
	t.Log(err)
}
