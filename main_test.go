// @Time : 2020/7/9 17:07
// @Author : 黑白配
// @File : main_test.go
// @PackageName:WXPay
// @Description:

package WXPay

import (
	"github.com/wleven/wxpay/global"
	"github.com/wleven/wxpay/src/V2"
	"github.com/wleven/wxpay/src/V3/smartGuide"
	"testing"
)

func TestInit(t *testing.T) {
	api := Init(global.V2)
	if data, err := api.V3.SmartGuide.Query(smartGuide.Query{StoreID: 20774227}); err == nil {
		t.Log(data)
	} else {
		t.Error(err)
	}

	if data, err := api.V2.WxAppPay(V2.UnifiedOrder{
		Attach:         "支付测试",
		OutTradeNo:     "11111111111115 ",
		TotalFee:       1,
		SpbillCreateIP: "127.0.0.1",
		OpenID:         "owJNp5PDj8lja9S3m2l2M_jt3aHY",
		Receipt:        "Y",
		Body:           "测试",
		TradeType:      "JSAPI",
	}); err == nil {
		t.Log(data)
	} else {
		t.Error(err)
	}
}
