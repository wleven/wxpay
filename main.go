// @Time : 2020/7/9 16:08
// @Author : 黑白配
// @File : main.go
// @PackageName:WXPay
// @Description:微信支付

package WXPay

import (
	"github.com/wleven/wxpay/src/V2"
	"github.com/wleven/wxpay/src/V3"
	"github.com/wleven/wxpay/src/config"
	"github.com/wleven/wxpay/src/entity"
)

type WXPayApi struct {
	V2     *V2.WxPay         // V2接口
	V3     *V3.API           // V3接口
	Config *entity.PayConfig // 配置
}

func Init(params entity.PayConfig) (api WXPayApi) {

	api.V2 = V2.Init(params)

	api.V3 = V3.Init(&config.V3{
		MchID:     params.MchID,
		ClientKey: params.APIClientPath.Key,
		SerialNo:  params.SerialNo,
	})
	api.Config = &params

	return
}
