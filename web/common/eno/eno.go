package eno

import (
	"errors"
	"fmt"
	"sync"
)

var (
	// CodeMap asd
	CodeMap = &sync.Map{}
	// Success 成功
	Success = New(0, "Success")
)

// RCode 码
type RCode struct {
	Code    int
	Message string
}

// Error error 接口
func (rc RCode) Error() string {
	return fmt.Sprintf("code:%d, message:%s", rc.Code, rc.Message)
}

// New 添加一个新的code码
func New(code int, message string) *RCode {
	CodeMap.Store(code, message)
	return &RCode{code, message}
}

// ParseRCode sadf
func ParseRCode(err error) (*RCode, error) {
	rc, ok := err.(*RCode)
	if !ok {
		return nil, errors.New("Parse rcode error")
	}
	return rc, nil
}
