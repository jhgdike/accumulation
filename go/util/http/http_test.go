package http

import (
"testing"
)

func TestGet(t *testing.T) {
	_, _, err := Get("http://www.baidu.com", nil, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestPost(t *testing.T) {
	_, _, err := PostJson("http://www.baidu.com", nil, nil)
	if err != nil {
		t.Error(err)
	}
}
