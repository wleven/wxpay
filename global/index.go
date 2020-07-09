// @Time : 2020/7/9 14:23
// @Author : 黑白配
// @File : index.go
// @PackageName:global
// @Description:

package global

import (
	"github.com/wleven/wxpay/src/config"
	"github.com/wleven/wxpay/src/entity"
)

var V3 = config.V3{
	MchID:         "",
	ClientKeyPath: "",
	SerialNo:      "",
}

var V2 = entity.PayConfig{
	AppID:        "",
	MchID:        "",
	SubAppID:     "",
	SubMchID:     "",
	PayNotify:    "",
	RefundNotify: "",
	Secret:       "",
	APIClientPath: entity.APIClientPath{
		Cert: "",
		Key:  "",
		Root: "",
	},
	SerialNo: "",
}
