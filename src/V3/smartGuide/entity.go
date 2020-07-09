// @Time : 2020/7/8 14:38
// @Author : 黑白配
// @File : entity.go
// @PackageName:smartGuide
// @Description:entity

package smartGuide

import "github.com/wleven/wxpay/src/config"

// Register 注册服务人员参数
type Register struct {
	SubMchID string `json:"sub_mchid,omitempty"` // 子商户ID
	CorpID   string `json:"corpid,omitempty"`    // 企业ID
	StoreID  int    `json:"store_id,omitempty"`  // 门店ID
	Name     string `json:"name,omitempty"`      // 员工姓名
	UserId   string `json:"userid,omitempty"`    // 企业微信的员工ID
	Mobile   string `json:"mobile,omitempty"`    // 手机号码
	QRCode   string `json:"qr_code,omitempty"`   // 二维码
	Avatar   string `json:"avatar,omitempty"`    // 头像
}

// RegisterResult 注册服务人员返回数据
type RegisterResult struct {
	GuideID string `json:"guide_id,omitempty"` // 服务人员ID
	config.Result
}

// Assign 服务人员分配参数
type Assign struct {
	GuideID    string `json:"guide_id,omitempty"`     // 服务人员ID
	SubMchID   string `json:"sub_mchid,omitempty"`    // 子商户ID
	OutTradeNo string `json:"out_trade_no,omitempty"` // 商户订单号
}

// Query 查询参数
type Query struct {
	SubMchID string `json:"sub_mchid,omitempty"` // 子商户ID
	StoreID  int    `json:"store_id,omitempty"`  // 门店ID
	UserId   string `json:"userid,omitempty"`    // 企业微信的员工ID
	Mobile   string `json:"mobile,omitempty"`    // 手机号码
	WorkID   string `json:"work_id,omitempty"`   // 工号
	Limit    int    `json:"limit,omitempty"`     // 最大资源条数
	Offset   int    `json:"offset,omitempty"`    // 请求资源起始位置
}

// UserInfo 服务人员信息
type UserInfo struct {
	GuideID string `json:"guide_id,omitempty"` // 服务人员ID
	StoreID int    `json:"store_id,omitempty"` // 门店ID
	Name    string `json:"name,omitempty"`     // 员工姓名
	UserId  string `json:"userid,omitempty"`   // 企业微信的员工ID
	Mobile  string `json:"mobile,omitempty"`   // 手机号码
	WorkID  string `json:"work_id,omitempty"`  // 工号
}

// QueryResult 查询结果
type QueryResult struct {
	Data       []UserInfo `json:"data,omitempty"`        // 服务人员列表
	TotalCount int        `json:"total_count,omitempty"` // 服务人员数量
	Limit      int        `json:"limit,omitempty"`       // 最大资源条数
	Offset     int        `json:"offset,omitempty"`      // 请求资源起始位置
	config.Result
}
