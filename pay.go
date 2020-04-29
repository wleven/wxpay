package pay

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/wleven/wxpay/utils"
)

// WxPay 微信支付
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

// 网络请求
func (c WxPay) request(url string, body io.Reader, cert bool) (map[string]string, error) {

	var client http.Client

	if cert {
		if err := c.checkClient(); err != nil {
			return nil, err
		}
		// 微信提供的API证书,证书和证书密钥 .pem格式
		certs, _ := tls.LoadX509KeyPair(c.APIClientPath.Cert, c.APIClientPath.Key)
		// 微信支付HTTPS服务器证书的根证书  .pem格式
		rootCa, _ := ioutil.ReadFile(c.APIClientPath.Root)

		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(rootCa)

		client = http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      pool,
				Certificates: []tls.Certificate{certs},
			},
		}}
	}

	if err := c.checkConfig(); err != nil {
		return nil, err
	}

	resp, err := client.Post(url, "", body)
	if err == nil {
		defer resp.Body.Close()
		b, _ := ioutil.ReadAll(resp.Body)
		var result PublicResponse
		_ = xml.Unmarshal(b, &result)
		err := result.ResultCheck()
		if err == nil {
			return utils.XML2MAP(b), nil
		}
		return nil, err
	}
	return nil, err
}

// 公共参数
func (c WxPay) publicParams() (m map[string]interface{}) {
	m = make(map[string]interface{})
	m["appid"] = c.AppID
	m["mch_id"] = c.MchID
	m["nonce_str"] = utils.RandomStr()
	if c.SubMchID != "" {
		m["sub_mch_id"] = c.SubMchID
	}
	if c.SubAppID != "" {
		m["sub_appid"] = c.SubAppID
	}
	m["sign_type"] = "HMAC-SHA256"
	return
}

// 检查支付配置
func (c WxPay) checkConfig() (err error) {

	if c.AppID == "" {
		err = errors.New("AppID 不能为空")
	} else if c.MchID == "" {
		err = errors.New("MchID 不能为空")
	} else if c.Secret == "" {
		err = errors.New("Secret 不能为空")
	} else {
		err = nil
	}
	return
}

// 检查支付证书
func (c WxPay) checkClient() (err error) {
	if c.APIClientPath.Cert == "" {
		err = errors.New("APIClientPath.Cert 不能为空")
	} else if c.APIClientPath.Key == "" {
		err = errors.New("APIClientPath.Key 不能为空")
	} else if c.APIClientPath.Root == "" {
		err = errors.New("APIClientPath.Key 不能为空")
	} else {
		err = nil
	}
	return
}

// UnifiedOrder 统一下单接口
func (c WxPay) UnifiedOrder(params UnifiedOrder) (map[string]string, error) {

	m := utils.MAPMerge(utils.Struct2Map(params), c.publicParams())

	m["notify_url"] = c.PayNotify
	m["sign"] = utils.Sign(m, c.Secret)

	return c.request("https://api.mch.weixin.qq.com/pay/unifiedorder", strings.NewReader(utils.MAP2XML(m)), false)
}

// WxAppPay 小程序下单接口
func (c WxPay) WxAppPay(params UnifiedOrder) (map[string]interface{}, error) {
	params.TradeType = "JSAPI"
	m, err := c.UnifiedOrder(params)
	if err == nil {
		result := make(map[string]interface{})
		result["appId"] = m["appid"]
		result["nonceStr"] = m["nonce_str"]
		result["package"] = "prepay_id=" + m["prepay_id"]
		result["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)
		result["signType"] = "HMAC-SHA256"
		result["paySign"] = utils.Sign(result, c.Secret)
		return result, err
	}
	return nil, err
}

// Micropay 付款码支付
func (c WxPay) Micropay(params Micropay) (map[string]string, error) {
	m := utils.MAPMerge(utils.Struct2Map(params), c.publicParams())
	m["sign"] = utils.Sign(m, c.Secret)
	return c.request("https://api.mch.weixin.qq.com/pay/micropay", strings.NewReader(utils.MAP2XML(m)), false)
}

// CloseOrder 关闭订单
func (c WxPay) CloseOrder(outTradeNo string) (map[string]string, error) {

	m := c.publicParams()

	m["out_trade_no"] = outTradeNo
	m["sign"] = utils.Sign(m, c.Secret)

	return c.request("https://api.mch.weixin.qq.com/pay/closeorder", strings.NewReader(utils.MAP2XML(m)), false)
}

// ReverseOrder 撤销订单
func (c WxPay) ReverseOrder(params ReverseOrder) (map[string]string, error) {
	m := utils.MAPMerge(utils.Struct2Map(params), c.publicParams())
	m["sign"] = utils.Sign(m, c.Secret)
	return c.request("https://api.mch.weixin.qq.com/secapi/pay/reverse", strings.NewReader(utils.MAP2XML(m)), true)
}

// OrderQuery 查询订单
func (c WxPay) OrderQuery(params OrderQuery) (map[string]string, error) {
	m := utils.MAPMerge(utils.Struct2Map(params), c.publicParams())
	m["sign"] = utils.Sign(m, c.Secret)
	return c.request("https://api.mch.weixin.qq.com/pay/orderquery", strings.NewReader(utils.MAP2XML(m)), false)
}

// Refund 申请退款
func (c WxPay) Refund(params Refund) (map[string]string, error) {
	m := utils.MAPMerge(utils.Struct2Map(params), c.publicParams())
	m["notify_url"] = c.RefundNotify
	m["sign"] = utils.Sign(m, c.Secret)
	return c.request("https://api.mch.weixin.qq.com/secapi/pay/refund", strings.NewReader(utils.MAP2XML(m)), true)
}

// RefundQuery 查询退款
func (c WxPay) RefundQuery(params RefundQuery) (map[string]string, error) {
	m := utils.MAPMerge(utils.Struct2Map(params), c.publicParams())
	m["sign"] = utils.Sign(m, c.Secret)
	return c.request("https://api.mch.weixin.qq.com/pay/refundquery", strings.NewReader(utils.MAP2XML(m)), false)
}

// ProfitSharingAddReceiver 添加分账接受方
func (c WxPay) ProfitSharingAddReceiver(params Receiver) (map[string]string, error) {
	m := c.publicParams()
	b, _ := json.Marshal(params)
	m["receiver"] = string(b)
	m["sign"] = utils.Sign(m, c.Secret)
	return c.request("https://api.mch.weixin.qq.com/pay/profitsharingaddreceiver", strings.NewReader(utils.MAP2XML(m)), false)
}

// ProfitSharingRemoveReceiver 删除分账接收方
func (c WxPay) ProfitSharingRemoveReceiver(params Receiver) (map[string]string, error) {
	m := c.publicParams()
	b, _ := json.Marshal(params)
	m["receiver"] = string(b)
	m["sign"] = utils.Sign(m, c.Secret)
	return c.request("https://api.mch.weixin.qq.com/pay/profitsharingremovereceiver", strings.NewReader(utils.MAP2XML(m)), false)
}

// ProfitSharingFinish 完成分账
func (c WxPay) ProfitSharingFinish(params ProfitSharingFinish) (map[string]string, error) {
	m := utils.MAPMerge(utils.Struct2Map(params), c.publicParams())
	m["sign"] = utils.Sign(m, c.Secret)
	return c.request("https://api.mch.weixin.qq.com/secapi/pay/profitsharingfinish", strings.NewReader(utils.MAP2XML(m)), true)
}

// ProfitSharing 开始分账 默认为单次分账 option=multi为多次分账
func (c WxPay) ProfitSharing(params ProfitSharing, option string) (map[string]string, error) {
	m := utils.MAPMerge(utils.Struct2Map(params), c.publicParams())
	b, _ := json.Marshal(params.Receivers)
	m["receivers"] = string(b)
	m["sign"] = utils.Sign(m, c.Secret)
	if option == "multi" {
		return c.request("https://api.mch.weixin.qq.com/secapi/pay/multiprofitsharing", strings.NewReader(utils.MAP2XML(m)), true)
	}
	return c.request("https://api.mch.weixin.qq.com/secapi/pay/profitsharing", strings.NewReader(utils.MAP2XML(m)), true)
}

// ProfitSharingQuery 查询分账结果
func (c WxPay) ProfitSharingQuery(params ProfitSharingQuery) (map[string]string, error) {
	m := utils.MAPMerge(utils.Struct2Map(params), c.publicParams())
	delete(m, "appid")
	delete(m, "sub_appid")

	m["sign"] = utils.Sign(m, c.Secret)
	return c.request("https://api.mch.weixin.qq.com/pay/profitsharingquery", strings.NewReader(utils.MAP2XML(m)), false)
}

// ProfitSharingReturn 分账回退
func (c WxPay) ProfitSharingReturn(params ProfitSharingReturn) (map[string]string, error) {
	m := utils.MAPMerge(utils.Struct2Map(params), c.publicParams())
	m["return_account_type"] = "MERCHANT_ID "
	m["sign"] = utils.Sign(m, c.Secret)
	return c.request("https://api.mch.weixin.qq.com/secapi/pay/profitsharingreturn", strings.NewReader(utils.MAP2XML(m)), true)
}

// ProfitSharingReturnQuery 分账回退结果查询
func (c WxPay) ProfitSharingReturnQuery(params ProfitSharingReturnQuery) (map[string]string, error) {
	m := utils.MAPMerge(utils.Struct2Map(params), c.publicParams())
	m["sign"] = utils.Sign(m, c.Secret)
	return c.request("https://api.mch.weixin.qq.com/pay/profitsharingreturnquery", strings.NewReader(utils.MAP2XML(m)), false)
}

// NotifyFormat 通知数据解析
func NotifyFormat(data string) map[string]string {
	return utils.XML2MAP([]byte(data))
}
