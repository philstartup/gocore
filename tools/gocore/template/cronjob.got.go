// Code generated by hero.
// source: /Users/liuguoqiang/Desktop/go/mod/gocore/tools/gocore/template/cronjob.got
// DO NOT EDIT!
package template

import (
	"bytes"

	"github.com/shiyanhui/hero"
)

func FromCronJob(cron string, buffer *bytes.Buffer) {
	buffer.WriteString(`
package cronjob
// `)
	hero.EscapeHTML(cron, buffer)
	buffer.WriteString(`
func `)
	hero.EscapeHTML(cron, buffer)
	buffer.WriteString(`() {
}`)

}
