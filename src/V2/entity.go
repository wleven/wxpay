package V2

import "errors"

// APIClientPath 微信支付API证书
type APIClientPath struct {
	Cert string // 证书路径
	Key  string // 证书密钥路径
	Root string // 根证书路径
}

// UnifiedOrder 统一下单参数
type UnifiedOrder struct {
	DeviceInfo     string              `json:"device_info,omitempty"`      // 自定义参数，可以为终端设备号(门店号或收银设备ID)
	Body           string              `json:"body,omitempty"`             // 商品简单描述，该字段请按照规范传递
	Detail         *UnifiedOrderDetail `json:"detail,omitempty"`           // 商品详情
	Attach         string              `json:"attach,omitempty"`           // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
	OutTradeNo     string              `json:"out_trade_no,omitempty"`     // 商户系统内部订单号
	FeeType        string              `json:"fee_type,omitempty"`         // 币种,默认人民币：CNY
	TotalFee       int                 `json:"total_fee,omitempty"`        // 订单总金额，单位为分
	SpbillCreateIP string              `json:"spbill_create_ip,omitempty"` // 支持IPV4和IPV6两种格式的IP地址。用户的客户端IP
	TimeStart      string              `json:"time_start,omitempty"`       // 订单生成时间
	TimeExpire     string              `json:"time_expire,omitempty"`      // 订单失效时间
	GoodsTag       string              `json:"goods_tag,omitempty"`        // 订单优惠标记
	NotifyURL      string              `json:"notify_url,omitempty"`       // 支付结果通知的回调地址
	TradeType      string              `json:"trade_type,omitempty"`       // JSAPI/NATIVE/APP
	ProductID      string              `json:"product_id,omitempty"`       // 商品ID trade_type=NATIVE时，此参数必传。
	LimitPay       string              `json:"limit_pay,omitempty"`        // no_credit--可限制用户不能使用信用卡支付
	OpenID         string              `json:"openid,omitempty"`           // 微信用户标识 trade_type=JSAPI时，此参数必传
	SubOpenid      string              `json:"sub_openid,omitempty"`       // 用户在子商户appid下的唯一标识。openid和sub_openid可以选传其中之一，如果选择传sub_openid,则必须传sub_appid
	Receipt        string              `json:"receipt,omitempty"`          // Y/N 是否开启电子发票入口
	SceneInfo      *SceneInfo          `json:"scene_info,omitempty"`       // 场景信息
	ProfitSharing  string              `json:"profit_sharing,omitempty"`   // 是否分账 Y/N
}

// Micropay 付款码支付
type Micropay struct {
	DeviceInfo     string              `json:"device_info,omitempty"`      // 自定义参数，可以为终端设备号(门店号或收银设备ID)
	Body           string              `json:"body,omitempty"`             // 商品简单描述，该字段请按照规范传递
	Detail         *UnifiedOrderDetail `json:"detail,omitempty"`           // 商品详情
	Attach         string              `json:"attach,omitempty"`           // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
	OutTradeNo     string              `json:"out_trade_no,omitempty"`     // 商户系统内部订单号
	FeeType        string              `json:"fee_type,omitempty"`         // 币种,默认人民币：CNY
	TotalFee       int                 `json:"total_fee,omitempty"`        // 订单总金额，单位为分
	SpbillCreateIP string              `json:"spbill_create_ip,omitempty"` // 支持IPV4和IPV6两种格式的IP地址。用户的客户端IP
	TimeStart      string              `json:"time_start,omitempty"`       // 订单生成时间
	TimeExpire     string              `json:"time_expire,omitempty"`      // 订单失效时间
	GoodsTag       string              `json:"goods_tag,omitempty"`        // 订单优惠标记
	LimitPay       string              `json:"limit_pay,omitempty"`        // no_credit--可限制用户不能使用信用卡支付
	SubOpenid      string              `json:"sub_openid,omitempty"`       // 用户在子商户appid下的唯一标识。openid和sub_openid可以选传其中之一，如果选择传sub_openid,则必须传sub_appid
	Receipt        string              `json:"receipt,omitempty"`          // Y/N 是否开启电子发票入口
	SceneInfo      *SceneInfo          `json:"scene_info,omitempty"`       // 场景信息
	ProfitSharing  string              `json:"profit_sharing,omitempty"`   // 是否分账 Y/N
	AuthCode       string              `json:"auth_code,omitempty"`        // 扫码支付付款码，设备读取用户微信中的条码或者二维码信息
}

// UnifiedOrderDetail 商品详情
type UnifiedOrderDetail struct {
	CostPrice   int           `json:"cost_price,omitempty"`   // 订单原价
	ReceiptID   string        `json:"receipt_id,omitempty"`   // 小票ID
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"` // 商品详情
}

// GoodsDetail 商品详情
type GoodsDetail struct {
	GoodsID      string `json:"goods_id,omitempty"`       // 商品ID
	WxpayGoodsID string `json:"wxpay_goods_id,omitempty"` // 微信支付定义的统一商品编号（没有可不传）
	GoodsName    string `json:"goods_name,omitempty"`     // 商品的实际名称
	Quantity     int    `json:"quantity,omitempty"`       // 用户购买的数量
	Price        int    `json:"price,omitempty"`          // 单位为：分。如果商户有优惠，需传输商户优惠后的单价
}

// SceneInfo 场景信息
type SceneInfo struct {
	ID       string `json:"id,omitempty"`        // 门店编号，由商户自定义
	Name     string `json:"name,omitempty"`      // 门店名称 ，由商户自定义
	AreaCode string `json:"area_code,omitempty"` // 门店所在地行政区划码
	Address  string `json:"address,omitempty"`   // 门店详细地址 ，由商户自定义
}

// ResultCheck 返回结果检查接口
type ResultCheck interface {
	ResultCheck() error
}

// PublicResponse 通用返回
type PublicResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`   // 此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`     // 当return_code为FAIL时返回信息为错误原因
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`   // SUCCESS/FAIL
	ResultMsg  string `xml:"result_msg,omitempty" json:"result_msg,omitempty"`     // 对于业务执行的详细描述
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`         // 当result_code为FAIL时返回错误代码
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"` // 当result_code为FAIL时返回错误描述
}

// ResultCheck 检查是否返回成功
func (m PublicResponse) ResultCheck() error {
	if m.ReturnCode == "FAIL" {
		return errors.New(m.ReturnMsg)
	} else if m.ResultCode == "FAIL" {
		return errors.New(m.ErrCodeDes)
	}
	return nil
}

// CDATA xml
type CDATA struct {
	Text string `xml:",cdata"`
}

// ReverseOrder 撤销订单
type ReverseOrder struct {
	TransactionID string `json:"transaction_id,omitempty"` // 微信支付ID 2选1
	OutTradeNo    string `json:"out_trade_no,omitempty"`   // 商户订单ID 2选1
}

// OrderQuery 查询订单参数
type OrderQuery struct {
	TransactionID string `json:"transaction_id,omitempty"` // 微信支付ID
	OutTradeNo    string `json:"out_trade_no,omitempty"`   // 商户订单ID
}

// Refund 退款参数
type Refund struct {
	TransactionID string `json:"transaction_id,omitempty"`  // 微信支付ID
	OutTradeNo    string `json:"out_trade_no,omitempty"`    // 商户订单ID
	OutRefundNo   string `json:"out_refund_no,omitempty"`   // 商户系统内部的退款单号
	TotalFee      int    `json:"total_fee,omitempty"`       // 订单总金额，单位为分
	RefundFee     int    `json:"refund_fee,omitempty"`      // 退款总金额，订单总金额
	RefundFeeType string `json:"refund_fee_type,omitempty"` // 退款货币类型，需与支付一致，或者不填。
	RefundDesc    string `json:"refund_desc,omitempty"`     // 若商户传入，会在下发给用户的退款消息中体现退款原因
	RefundAccount string `json:"refund_account,omitempty"`  // 退款资金来源 仅针对老资金流商户使用
}

// RefundQuery 退款查询参数
type RefundQuery struct {
	TransactionID string `json:"transaction_id,omitempty"` // 微信支付ID 四选一
	OutTradeNo    string `json:"out_trade_no,omitempty"`   // 商户订单ID 四选一
	OutRefundNo   string `json:"out_refund_no,omitempty"`  // 商户系统内部的退款单号 四选一
	RefundID      string `json:"refund_id,omitempty"`      // 微信退款单号 四选一
	Offset        int    `json:"offset,omitempty"`         // 偏移量，当部分退款次数超过10次时可使用，表示返回的查询结果从这个偏移量开始取记录
}

// Receiver 分账接收方
type Receiver struct {
	Type           string `json:"type,omitempty"`            // 账号类型
	Account        string `json:"account,omitempty"`         // 接收方账号
	Name           string `json:"name,omitempty"`            // 商户全称
	CustomRelation string `json:"custom_relation,omitempty"` // 子商户与接收方具体的关系，本字段最多10个字
	RelationType   string `json:"relation_type,omitempty"`   // 分账关系
	Amount         int    `json:"amount,omitempty"`          // 分账金额
	Description    string `json:"description,omitempty"`     // 分账描述
}

// ProfitSharingFinish 分账完成参数
type ProfitSharingFinish struct {
	TransactionID string `json:"transaction_id,omitempty"` // 微信支付ID
	OutOrderNo    string `json:"out_order_no,omitempty"`   // 商户订单ID
	Description   string `json:"description,omitempty"`    // 分账完结的原因描述
}

// ProfitSharing 单次分账参数
type ProfitSharing struct {
	TransactionID string     `json:"transaction_id,omitempty"` // 微信支付ID
	OutOrderNo    string     `json:"out_order_no,omitempty"`   // 商户订单ID
	Receivers     []Receiver `json:"receivers,omitempty"`      // 接收方列表
}

// ProfitSharingQuery 分账结果查询参数
type ProfitSharingQuery struct {
	TransactionID string `json:"transaction_id,omitempty"` // 微信支付ID
	OutOrderNo    string `json:"out_order_no,omitempty"`   // 商户订单ID
}

// ProfitSharingReturn 分账回退参数
type ProfitSharingReturn struct {
	OrderID       string `json:"order_id,omitempty"`       // 原发起分账请求时，微信返回的微信分账单号，与商户分账单号一一对应 2选1
	OutOrderNo    string `json:"out_order_no,omitempty"`   // 原发起分账请求时使用的商户后台系统的分账单号 2选1
	OutReturnNo   string `json:"out_return_no,omitempty"`  // 商户在自己后台生成的一个新的回退单号
	ReturnAccount string `json:"return_account,omitempty"` // 回退方类型是MERCHANT_ID时，填写商户ID
	ReturnAmount  int    `json:"return_amount,omitempty"`  // 回退金额
	Description   string `json:"description,omitempty"`    // 分账回退的原因描述
}

// ProfitSharingReturnQuery 分账回退查询参数
type ProfitSharingReturnQuery struct {
	OrderID     string `json:"order_id,omitempty"`      // 原发起分账请求时，微信返回的微信分账单号，与商户分账单号一一对应 2选1
	OutOrderNo  string `json:"out_order_no,omitempty"`  // 原发起分账请求时使用的商户后台系统的分账单号 2选1
	OutReturnNo string `json:"out_return_no,omitempty"` // 商户系统内部的回退单号
}

// Transfers 付款到零钱
type Transfers struct {
	PartnerTradeNo string `json:"partner_trade_no"` // 商户订单号
	OpenID         string `json:"openid"`           // OPENID
	CheckName      string `json:"check_name"`       // NO_CHECK：不校验真实姓名  FORCE_CHECK：强校验真实姓名
	ReUserName     string `json:"re_user_name"`     // 收款用户真实姓名。如果check_name设置为FORCE_CHECK，则必填用户真实姓名。如需电子回单，需要传入收款用户姓名
	Amount         int    `json:"amount"`           // 企业付款金额，单位为分
	Desc           string `json:"desc"`             // 企业付款备注，必填。注意：备注中的敏感词会被转成字符*
}
