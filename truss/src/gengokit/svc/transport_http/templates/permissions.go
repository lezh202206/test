package templates

const Permissions = `
package transport_http

import (
	"gitee.com/zhubaoe-go/jordan/auth"
	"gitee.com/zhubaoe-go/jordan/auth/permit_type"
)

var Permissions = map[string]auth.Permission{
	"/echo": {
		Path: "/{{ToLower .Service.Name}}/echo",
		Type: permit_type.None,
	},
}
`
