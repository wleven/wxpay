// @Time : 2020/7/9 16:51
// @Author : 黑白配
// @File : index
// @PackageName:entity
// @Description:

package entity

type PayConfig struct {
	AppID         string        // 商户/服务商 AppId(公众号/小程序)
	MchID         string        // 商户/服务商 商户号
	SubAppID      string        // 子商户公众号ID
	SubMchID      string        // 子商户商户号
	PayNotify     string        // 支付结果回调地址
	RefundNotify  string        // 退款结果回调地址
	Secret        string        // 微信支付密钥
	APIClientPath APIClientPath // API证书路径,使用V3接口必传
	SerialNo      string        // 证书编号,使用V3接口必传
}

// APIClientPath 微信支付API证书
type APIClientPath struct {
	Cert string // 证书路径
	Key  string // 私钥证书路径,使用V3接口必传
	Root string // 根证书路径
}
