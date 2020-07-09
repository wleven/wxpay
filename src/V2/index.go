// @Time : 2020/7/9 16:03
// @Author : 黑白配
// @File : index.go
// @PackageName:V2
// @Description:V2版本接口

package V2

import "github.com/wleven/wxpay/src/entity"

func Init(config entity.PayConfig) *WxPay {
	return &WxPay{config: config}
}
