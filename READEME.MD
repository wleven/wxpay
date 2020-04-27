# 微信支付 SDK For Golang

### 支持微信商户支付/商户分账/服务商支付/服务商分账

- [安装包](#安装包)
- [支付环境初始化](#支付环境初始化)
- 支付接口
  - [统一下单](#统一下单)
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
go get -u  https://github.com/wleven/wxpay
```

## 支付环境初始化

```
var pay = WxPay{
    ...
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
	APIClientPath APIClientPath // API证书路径
}

// APIClientPath 微信支付API证书
type APIClientPath struct {
	Cert string // 证书路径
	Key  string // 证书密钥路径
	Root string // 根证书路径
}

```

## 统一下单

```
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
		log.Println(err)
	}
```

## 付款码支付

```

```

## 关闭订单

```

```

## 撤销订单

```

```

## 查询订单

```

```

## 申请退款

```

```

## 查询退款

```

```

## 添加分账接收方

```

```

## 删除分账接收方

```

```

## 发起分账

```

```

## 完成分账

```

```

## 查询分账结果

```

```

## 分账回退

```

```

## 分账回退结果查询

```

```
