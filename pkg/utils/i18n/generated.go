package i18n

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// initEnUS will init en_US support.
func initEnUS(tag language.Tag) {
	_ = message.SetString(tag, "findv_description", "A simple tool for scanning artifact vulnerabilities.")
}

// initZhCN will init zh_CN support.
func initZhCN(tag language.Tag) {
	_ = message.SetString(tag, "findv_description", "一个简单的用于扫描制品漏洞的工具。")
}
