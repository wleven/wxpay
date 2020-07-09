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
	V2 *V2.WxPay
	V3 *V3.API
}

func Init(params entity.PayConfig) (api WXPayApi) {

	api.V2 = V2.Init(params)

	api.V3 = V3.Init(&config.V3{
		MchID:         params.MchID,
		ClientKeyPath: params.APIClientPath.Key,
		SerialNo:      params.SerialNo,
	})

	return
}
