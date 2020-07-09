package V2

import (
	"github.com/wleven/wxpay/global"
	"log"
	"testing"
)

var pay = WxPay{
	global.V2,
}

func TestWxPay_UnifiedOrder(t *testing.T) {
	if result, err := pay.UnifiedOrder(UnifiedOrder{
		Attach:         "支付测试",
		OutTradeNo:     "11111111111114",
		TotalFee:       1,
		SpbillCreateIP: "127.0.0.1",
		OpenID:         "owJNp5PDj8lja9S3m2l2M_jt3aHY",
		Receipt:        "Y",
		Body:           "测试",
		TradeType:      "JSAPI",
	}); err == nil {
		log.Println(result)
	} else {
		t.Error(err)
	}
}

func TestWxPay_WxAppPay(t *testing.T) {
	if result, err := pay.WxAppPay(UnifiedOrder{
		Attach:         "支付测试",
		OutTradeNo:     "11111111111115 ",
		TotalFee:       1,
		SpbillCreateIP: "127.0.0.1",
		OpenID:         "owJNp5PDj8lja9S3m2l2M_jt3aHY",
		Receipt:        "Y",
		Body:           "测试",
		TradeType:      "JSAPI",
	}); err == nil {
		log.Println(result)
	} else {
		t.Error(err)
	}
}

func TestWxPay_Micropay(t *testing.T) {
	if result, err := pay.Micropay(Micropay{
		Attach:         "支付测试",
		OutTradeNo:     "11111111111115",
		TotalFee:       1,
		SpbillCreateIP: "127.0.0.1",
		Receipt:        "Y",
		Body:           "测试",
		AuthCode:       "12312312312",
	}); err == nil {
		log.Println(result)
	} else {
		t.Error(err)
	}
}

func TestWxPay_CloseOrder(t *testing.T) {
	if result, err := pay.CloseOrder("11111111111112"); err == nil {
		log.Println(result)
	} else {
		t.Error(err)
	}
}

func TestWxPay_ReverseOrder(t *testing.T) {
	if result, err := pay.ReverseOrder(ReverseOrder{
		OutTradeNo: "11111111111112",
	}); err == nil {
		log.Println(result)
	} else {
		t.Error(err)
	}
}

func TestWxPay_OrderQuery(t *testing.T) {
	if result, err := pay.OrderQuery(OrderQuery{
		OutTradeNo: "674BB66E408A6931788347BF25E9BCAA",
	}); err == nil {
		log.Println(result)
	} else {
		t.Error(err)
	}
}

func TestWxPay_Refund(t *testing.T) {
	if result, err := pay.Refund(Refund{
		OutTradeNo:  "BC9AA9B9E43C13E7932CB3B181468A4F",
		TotalFee:    990,
		RefundFee:   990,
		OutRefundNo: "BC9AA9B9E43C13E7932CB3B181468A4F",
	}); err == nil {
		log.Println(result)
	} else {
		t.Error(err)
	}
}

func TestWxPay_RefundQuery(t *testing.T) {
	if result, err := pay.RefundQuery(RefundQuery{
		OutTradeNo: "11111111111113",
	}); err == nil {
		log.Println(result)
	} else {
		t.Error(err)
	}
}

func TestWxPay_ProfitSharingAddReceiver(t *testing.T) {
	if result, err := pay.ProfitSharingAddReceiver(Receiver{
		Type:         "PERSONAL_OPENID",
		Account:      "owJNp5M6Trxg5qBjh8KTPnTm65Sg",
		RelationType: "DISTRIBUTOR",
	}); err == nil {
		log.Println(result)
	} else {
		t.Error(err)
	}
}

func TestWxPay_ProfitSharingRemoveReceiver(t *testing.T) {
	if result, err := pay.ProfitSharingRemoveReceiver(Receiver{
		Type:    "PERSONAL_OPENID",
		Account: "owJNp5PDj8lja9S3m2l2M_jt3aHY",
	}); err == nil {
		log.Println(result)
	} else {
		t.Error(err)
	}
}

func TestWxPay_ProfitSharingFinish(t *testing.T) {
	if result, err := pay.ProfitSharingFinish(ProfitSharingFinish{
		TransactionID: "11111",
		OutOrderNo:    "111111",
		Description:   "分账完成",
	}); err == nil {
		log.Println(result)
	} else {
		t.Error(err)
	}
}

func TestWxPay_ProfitSharing(t *testing.T) {
	if result, err := pay.ProfitSharing(ProfitSharing{
		TransactionID: "11111",
		OutOrderNo:    "111111",
		Receivers:     []Receiver{{Type: "PERSONAL_OPENID", Account: "owJNp5PDj8lja9S3m2l2M_jt3aHY", Amount: 1, Description: "个人分成"}},
	}, "multi"); err == nil {
		log.Println(result)
	} else {
		t.Error(err)
	}
}

func TestWxPay_ProfitSharingQuery(t *testing.T) {
	if result, err := pay.ProfitSharingQuery(ProfitSharingQuery{
		TransactionID: "11111",
		OutOrderNo:    "111111",
	}); err == nil {
		log.Println(result)
	} else {
		t.Error(err)
	}
}

func TestWxPay_ProfitSharingReturn(t *testing.T) {
	if result, err := pay.ProfitSharingReturn(ProfitSharingReturn{
		OutOrderNo:    "1111111",
		OutReturnNo:   "1111112",
		ReturnAccount: "12312312",
		ReturnAmount:  100,
		Description:   "回退",
	}); err == nil {
		log.Println(result)
	} else {
		t.Error(err)
	}
}

func TestWxPay_ProfitSharingReturnQuery(t *testing.T) {
	if result, err := pay.ProfitSharingReturnQuery(ProfitSharingReturnQuery{
		OutOrderNo:  "1111111",
		OutReturnNo: "1111112",
	}); err == nil {
		log.Println(result)
	} else {
		t.Error(err)
	}
}

func TestNotifyDataFormat(t *testing.T) {
	x := `<xml>
	<appid><![CDATA[wx2421b1c4370ec43b]]></appid>
	<attach><![CDATA[支付测试]]></attach>
	<bank_type><![CDATA[CFT]]></bank_type>
	<fee_type><![CDATA[CNY]]></fee_type>
	<is_subscribe><![CDATA[Y]]></is_subscribe>
	<mch_id><![CDATA[10000100]]></mch_id>
	<nonce_str><![CDATA[5d2b6c2a8db53831f7eda20af46e531c]]></nonce_str>
	<openid><![CDATA[oUpF8uMEb4qRXf22hE3X68TekukE]]></openid>
	<out_trade_no><![CDATA[1409811653]]></out_trade_no>
	<result_code><![CDATA[SUCCESS]]></result_code>
	<return_code><![CDATA[SUCCESS]]></return_code>
	<sign><![CDATA[B552ED6B279343CB493C5DD0D78AB241]]></sign>
	<time_end><![CDATA[20140903131540]]></time_end>
	<total_fee>1</total_fee>
  <coupon_fee><![CDATA[10]]></coupon_fee>
  <coupon_count><![CDATA[1]]></coupon_count>
  <coupon_type><![CDATA[CASH]]></coupon_type>
  <coupon_id><![CDATA[10000]]></coupon_id>
  <coupon_fee><![CDATA[100]]></coupon_fee>
	<trade_type><![CDATA[JSAPI]]></trade_type>
	<transaction_id><![CDATA[1004400740201409030005092168]]></transaction_id>
  </xml>`
	m := NotifyFormat(x)

	log.Println(m)

}
