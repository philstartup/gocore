// Code generated by hero.
// source: /Users/SM0410/project/go/gocore/tools/gocore/template/err_code.got
// DO NOT EDIT!
package template

import "bytes"

func FromErrCode(buffer *bytes.Buffer) {
	buffer.WriteString(`
package errcode

import (
	"github.com/sunmi-OS/gocore/v2/api/ecode"
	"gorm.io/gorm"
)

var (
	ErrorNotFound = ecode.NewV2(50001, "record not found")
)`)

}
