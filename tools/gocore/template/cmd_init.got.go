// Code generated by hero.
// source: /Users/SM0286/code/core/gocore/tools/gocore/template/cmd_init.got
// DO NOT EDIT!
package template

import (
	"bytes"

	"github.com/shiyanhui/hero"
)

func FromCmdInit(name, pkgs, dbUpdate, initDb string, buffer *bytes.Buffer) {
	buffer.WriteString(`
package cmd

import (
	"log"

	`)
	hero.EscapeHTML(pkgs, buffer)
	buffer.WriteString(`
	"`)
	hero.EscapeHTML(name, buffer)
	buffer.WriteString(`/conf"

	"github.com/sunmi-OS/gocore/v2/db/gorm"
	"github.com/sunmi-OS/gocore/v2/conf/nacos"
	"github.com/sunmi-OS/gocore/v2/utils"
)

// initConf 初始化配置服务 （内部方法）
func initConf() {
	// 初始化Nacos配置
	conf.InitNacos(utils.GetRunTime())
	// 注册需要的配置
	nacos.ViperTomlHarder.SetDataIds("`)
	hero.EscapeHTML(name, buffer)
	buffer.WriteString(`", "mysql", "config", "redis")
	// 注册配置更新回调
	nacos.ViperTomlHarder.SetCallBackFunc("`)
	hero.EscapeHTML(name, buffer)
	buffer.WriteString(`", "mysql", func(namespace, group, dataId, data string) {
		`)
	hero.EscapeHTML(dbUpdate, buffer)
	buffer.WriteString(`
	})
	// 把Nacos的配置注册到Viper
	nacos.ViperTomlHarder.NacosToViper()
}

// initDB 初始化DB服务 （内部方法）
func initDB() {
	`)
	hero.EscapeHTML(initDb, buffer)
	buffer.WriteString(`
}`)

}
