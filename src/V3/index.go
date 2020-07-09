// @Time : 2020/7/8 16:04
// @Author : 黑白配
// @File : index.go
// @PackageName:V3
// @Description:V3版本接口

package V3

import (
	"github.com/wleven/wxpay/src/V3/smartGuide"
	"github.com/wleven/wxpay/src/config"
)

type API struct {
	SmartGuide *smartGuide.SmartGuide
}

func Init(params *config.V3) *API {
	return &API{SmartGuide: smartGuide.New(params)}
}
