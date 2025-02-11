// Code generated by hero.
// DO NOT EDIT!
package template

import "bytes"

func FromApiRoutes(name, routes string, buffer *bytes.Buffer) {
	buffer.WriteString(`
package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http/pprof"

	"`)
	buffer.WriteString(name)
	buffer.WriteString(`/api"
)

func Routes(router *gin.Engine) {
    `)
	buffer.WriteString(routes)
	buffer.WriteString(`
}
`)

}
