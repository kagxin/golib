package eno

import (
	"testing"
)

func Test_New(t *testing.T) {
	var (
		SUCESSS = New(0, "Success")
	)

	rCode, err := ParseRCode(SUCESSS)
	if err != nil {
		t.Fail()
	}
	if rCode.Code != 0 {
		t.Fail()
	}
	t.Logf("code:%d, message:%s", rCode.Code, rCode.Message)
}
