# 微信支付 SDK

- [x] V2版支付(商户/服务商)
- [x] V2版分账(商户/服务商)
- [x] V3版支付即服务

![go 1.14](https://img.shields.io/badge/go-1.14-green)
[![go.dev doc](https://img.shields.io/badge/go.dev-doc-green)](https://pkg.go.dev/github.com/wleven/wxpay)
[![GitHub license](https://img.shields.io/github/license/wleven/wxpay)](https://github.com/wleven/wxpay/blob/master/LICENSE)



- [安装包](#安装包)
- [查看文档](#查看文档)
- [V2版本下单接口](#V2版本下单接口)
- [V2版本分账接口](#V2版本分账接口)
- [V3版本支付即服务接口](#V3版本支付即服务接口)

## 安装包

```golang
go get -u  github.com/wleven/wxpay
```

## 查看文档

```golang
// 执行命令
godoc -http=:8888 -play
// 浏览器打开文档
http://127.0.0.1:8888/pkg/github.com/wleven/wxpay/
```

## V2版本下单接口

```golang
        config := entity.PayConfig{
 		// 传入支付初始化参数
 		// AppID         string        // 商户/服务商 AppId(公众号/小程序)
 		// MchID         string        // 商户/服务商 商户号
 		// SubAppID      string        // 子商户公众号ID
 		// SubMchID      string        // 子商户商户号
 		// PayNotify     string        // 支付结果回调地址
 		// RefundNotify  string        // 退款结果回调地址
 		// Secret        string        // 微信支付密钥
 		// APIClientPath APIClientPath // API证书路径,使用V3接口必传
 		// SerialNo      string        // 证书编号,使用V3接口必传
 	}

 	wxpay := WXPay.Init(config)
 
 	// 统一下单
 	if data, err := wxpay.V2.UnifiedOrder(V2.UnifiedOrder{/* 传入参数 */}); err == nil {
 	}
 	// 小程序支付
 	if data, err := wxpay.V2.WxAppPay(V2.UnifiedOrder{/* 传入参数 */}); err == nil {
 	}
 	// 付款码支付
 	if data, err := wxpay.V2.Micropay(V2.Micropay{/* 传入参数 */}); err == nil {
 	}
 	// 关闭订单
 	if data, err := wxpay.V2.CloseOrder("1111"); err == nil {
 	}
 	// 撤销订单
 	if data, err := wxpay.V2.ReverseOrder(V2.ReverseOrder{/* 传入参数 */}); err == nil {
 	}
 	// 查询订单
 	if data, err := wxpay.V2.OrderQuery(V2.OrderQuery{/* 传入参数 */}); err == nil {
 	}
 	// 申请退款
 	if data, err := wxpay.V2.Refund(V2.Refund{/* 传入参数 */}); err == nil {
 	}
 	// 查询退款
 	if data, err := wxpay.V2.RefundQuery(V2.RefundQuery{/* 传入参数 */}); err == nil {
 	}

```

## V2版本分账接口
```golang
        // 添加分账接收方
    	if data, err := wxpay.V2.ProfitSharingAddReceiver(V2.Receiver{/* 传入参数 */}); err == nil {
    	}
    	// 删除分账接收方
    	if data, err := wxpay.V2.ProfitSharingRemoveReceiver(V2.Receiver{/* 传入参数 */}); err == nil {
    	}
    	// 发起分账 第二个参数options为multi为多次分账  默认为单次
    	if data, err := wxpay.V2.ProfitSharing(V2.ProfitSharing{/* 传入参数 */},""); err == nil {
    	}
    	// 完成分账
    	if data, err := wxpay.V2.ProfitSharingFinish(V2.ProfitSharingFinish{/* 传入参数 */}); err == nil {
    	}
    	// 查询分账结果
    	if data, err := wxpay.V2.ProfitSharingQuery(V2.ProfitSharingQuery{/* 传入参数 */}); err == nil {
    	}
    	// 分账回退
    	if data, err := wxpay.V2.ProfitSharingReturn(V2.ProfitSharingReturn{/* 传入参数 */}); err == nil {
    	}
    	// 分账回退结果查询
    	if data, err := wxpay.V2.ProfitSharingReturnQuery(V2.ProfitSharingReturnQuery{/* 传入参数 */}); err == nil {
    	}

```

## V3版本支付即服务接口
```golang
    	// 注册服务人员
    	if data, err := wxpay.V3.SmartGuide.Register(smartGuide.Register{/* 传入参数 */}); err == nil {
    	}
    	// 分配服务人员
    	if data, err := wxpay.V3.SmartGuide.Assign(smartGuide.Assign{/* 传入参数 */}); err == nil {
    	}
    	// 查询服务人员
    	if data, err := wxpay.V3.SmartGuide.Query(smartGuide.Query{/* 传入参数 */}); err == nil {
    	}
```