// Code generated by hero.
// source: /Users/liuguoqiang/Desktop/go/mod/gocore/tools/gocore/template/conf_const.got
// DO NOT EDIT!
package template

import (
	"bytes"
	"strings"

	"github.com/shiyanhui/hero"
)

func FromConfConst(projectName string, buffer *bytes.Buffer) {
	buffer.WriteString(`
package conf

const (
	ProjectName    = "`)
	hero.EscapeHTML(projectName, buffer)
	buffer.WriteString(`"
	ProjectVersion = "v1.0.0"
	`)
	for _, v1 := range goCoreConfig.Config.CMysql {
		buffer.WriteString(`
		DB`)
		strings.Title(v1.Name)
		buffer.WriteString(` = "db`)
		strings.Title(v1.Name)
		buffer.WriteString(`"
	`)
	}
	for _, v1 := range goCoreConfig.Config.CRedis {
		for k2 := range v1.Index {
			buffer.WriteString(`
		DB`)
			hero.EscapeHTML(strings.Title(v1.Name)+strings.Title(k2), buffer)
			buffer.WriteString(`Redis = "`)
			hero.EscapeHTML(v1.Name, buffer)
			buffer.WriteString(`.`)
			hero.EscapeHTML(k2, buffer)
			buffer.WriteString(`"
	  `)
		}
	}
	buffer.WriteString(`
)`)

}
