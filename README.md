# 微信支付 SDK For Golang

![go 1.14](https://img.shields.io/badge/go-1.14-green)
[![go.dev doc](https://img.shields.io/badge/go.dev-doc-green)](https://pkg.go.dev/github.com/wleven/wxpay)
[![GitHub license](https://img.shields.io/github/license/wleven/wxpay)](https://github.com/wleven/wxpay/blob/master/LICENSE)

### 支持微信商户支付/商户分账/服务商支付/服务商分账

- [安装包](#安装包)
- [查看文档](#查看文档)
- [支付环境初始化](#支付环境初始化)
- 支付接口
  - [统一下单](#统一下单)
  - [小程序支付](#小程序支付)
  - [付款码支付](#付款码支付)
  - [关闭订单](#关闭订单)
  - [撤销订单](#撤销订单)
  - [查询订单](#查询订单)
  - [申请退款](#申请退款)
  - [查询退款](#查询退款)
- 分账接口
  - [添加分账接收方](#添加分账接收方)
  - [删除分账接收方](#删除分账接收方)
  - [发起单次分账](#发起分账)
  - [发起多次分账](#发起分账)
  - [完成分账](#完成分账)
  - [查询分账结果](#查询分账结果)
  - [分账回退](#分账回退)
  - [分账回退结果查询](#分账回退结果查询)

## 安装包

```
go get -u  github.com/wleven/wxpay
```

## 查看文档

```
// 执行命令
godoc -http=:8888 -play
// 浏览器打开文档
http://127.0.0.1:8888/pkg/github.com/wleven/wxpay/
```

## 支付环境初始化

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)
[![APIClientPath Doc](https://img.shields.io/badge/APIClientPath-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#APIClientPath)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}

	// WxPay 支付环境
	type WxPay struct {
		AppID         string        // 商户/服务商 AppId(公众号/小程序)
		MchID         string        // 商户/服务商 商户号
		SubAppID      string        // 子商户公众号ID
		SubMchID      string        // 子商户商户号
		PayNotify     string        // 支付结果回调地址
		RefundNotify  string        // 退款结果回调地址
		Secret        string        // 微信支付密钥
		APIClientPath pay.APIClientPath // API证书路径
	}

	// APIClientPath 微信支付API证书
	type APIClientPath struct {
		Cert string // 证书路径
		Key  string // 证书密钥路径
		Root string // 根证书路径
	}

```

## 统一下单

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)
[![UnifiedOrder Doc](https://img.shields.io/badge/UnifiedOrder-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#UnifiedOrder)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}
    if result, err := WxPay.UnifiedOrder(pay.UnifiedOrder{
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
		log.Println(err)
	}
```

## 小程序支付

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)
[![UnifiedOrder Doc](https://img.shields.io/badge/UnifiedOrder-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#UnifiedOrder)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}
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
		log.Println(err)
	}

```

## 付款码支付

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)
[![Micropay Doc](https://img.shields.io/badge/Micropay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#Micropay)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}
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
		log.Println(err)
	}

```

## 关闭订单

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}
	if result, err := WxPay.CloseOrder("11111111111112"); err == nil {
		log.Println(result)
	} else {
		log.Println(err)
	}

```

## 撤销订单

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)
[![ReverseOrder Doc](https://img.shields.io/badge/ReverseOrder-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#ReverseOrder)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}
	if result, err := pay.ReverseOrder(ReverseOrder{
		OutTradeNo: "11111111111112",
	}); err == nil {
		log.Println(result)
	} else {
		log.Println(err)
	}

```

## 查询订单

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)
[![OrderQuery Doc](https://img.shields.io/badge/OrderQuery-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#OrderQuery)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}
	if result, err := WxPay.OrderQuery(pay.OrderQuery{
		OutTradeNo: "11111111111113",
	}); err == nil {
		log.Println(result)
	} else {
		log.Println.Error(err)
	}
```

## 申请退款

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)
[![Refund Doc](https://img.shields.io/badge/Refund-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#Refund)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}
	if result, err := WxPay.Refund(pay.Refund{
		OutTradeNo:  "11111111111113",
		TotalFee:    100,
		RefundFee:   1,
		OutRefundNo: "11111111111113",
	}); err == nil {
		log.Println(result)
	} else {
		log.Println(err)
	}
```

## 查询退款

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)
[![RefundQuery Doc](https://img.shields.io/badge/RefundQuery-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#RefundQuery)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}
	if result, err := pay.RefundQuery(RefundQuery{
		OutTradeNo: "11111111111113",
	}); err == nil {
		log.Println(result)
	} else {
		log.Println(err)
	}

```

## 添加分账接收方

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)
[![Receiver Doc](https://img.shields.io/badge/Receiver-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#Receiver)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}
	if result, err := WxPay.ProfitSharingAddReceiver(pay.Receiver{
		Type:         "PERSONAL_OPENID",
		Account:      "owJNp5PDj8lja9S3m2l2M_jt3aHY",
		RelationType: "DISTRIBUTOR",
	}); err == nil {
		log.Println(result)
	} else {
		log.Println(err)
	}

```

## 删除分账接收方

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)
[![Receiver Doc](https://img.shields.io/badge/Receiver-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#Receiver)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}
    if result, err := WxPay.ProfitSharingRemoveReceiver(pay.Receiver{
		Type:    "PERSONAL_OPENID",
		Account: "owJNp5PDj8lja9S3m2l2M_jt3aHY",
	}); err == nil {
		log.Println(result)
	} else {
		log.Println(err)
	}

```

## 发起分账

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)
[![ProfitSharing Doc](https://img.shields.io/badge/ProfitSharing-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#ProfitSharing)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}
	// 第二个参数options为multi为多次分账  默认为单次
	if result, err := WxPay.ProfitSharing(pay.ProfitSharing{
		TransactionID: "11111",
		OutOrderNo:    "111111",
		Receivers:     []Receiver{{Type: "PERSONAL_OPENID", Account: "owJNp5PDj8lja9S3m2l2M_jt3aHY", Amount: 1, Description: "个人分成"}},
	}, "multi"); err == nil {
		log.Println(result)
	} else {
		log.Println(err)
	}

```

## 完成分账

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)
[![ProfitSharingFinish Doc](https://img.shields.io/badge/ProfitSharingFinish-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#ProfitSharingFinish)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}
	if result, err := WxPay.ProfitSharingFinish(pay.ProfitSharingFinish{
		TransactionID: "11111",
		OutOrderNo:    "111111",
		Description:   "分账完成",
	}); err == nil {
		log.Println(result)
	} else {
		log.Println(err)
	}

```

## 查询分账结果

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)
[![ProfitSharingQuery Doc](https://img.shields.io/badge/ProfitSharingQuery-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#ProfitSharingQuery)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}
	if result, err := WxPay.ProfitSharingQuery(pay.ProfitSharingQuery{
		TransactionID: "11111",
		OutOrderNo:    "111111",
	}); err == nil {
		log.Println(result)
	} else {
		log.Println(err)
	}

```

## 分账回退

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)
[![ProfitSharingReturn Doc](https://img.shields.io/badge/ProfitSharingReturn-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#ProfitSharingReturn)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}
	if result, err := WxPay.ProfitSharingReturn(pay.ProfitSharingReturn{
		OutOrderNo:    "1111111",
		OutReturnNo:   "1111112",
		ReturnAccount: "12312312",
		ReturnAmount:  100,
		Description:   "回退",
	}); err == nil {
		log.Println(result)
	} else {
		log.Println(err)
	}

```

## 分账回退结果查询

[![WxPay Doc](https://img.shields.io/badge/WxPay-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#WxPay)
[![ProfitSharingReturnQuery Doc](https://img.shields.io/badge/ProfitSharingReturnQuery-Struct_Doc-green)](https://pkg.go.dev/github.com/wleven/wxpay?tab=doc#ProfitSharingReturnQuery)

```
	var WxPay = pay.WxPay{
		// 支付初始化参数
	}
	if result, err := WxPay.ProfitSharingReturnQuery(pay.ProfitSharingReturnQuery{
		OutOrderNo:  "1111111",
		OutReturnNo: "1111112",
	}); err == nil {
		log.Println(result)
	} else {
		log.Println(err)
	}

```
