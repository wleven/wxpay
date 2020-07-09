// @Time : 2020/7/8 14:26
// @Author : 黑白配
// @File : index.go
// @PackageName:smartGuide
// @Description: 支付即服务

package smartGuide

import (
	"github.com/wleven/wxpay/src/config"
	"github.com/wleven/wxpay/utils"
)

func New(config *config.V3) *SmartGuide {
	return &SmartGuide{config: config}
}

type SmartGuide struct {
	config *config.V3
}

// Register 服务人员注册
//
// https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/smartguide/chapter3_1.shtml
func (m SmartGuide) Register(register Register) (result RegisterResult, err error) {
	err = m.config.Request("/v3/smartguide/guides", register, &result)
	return
}

// Assign 服务人员分配
//
// https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/smartguide/chapter3_1.shtml
func (m SmartGuide) Assign(assign Assign) (result interface{}, err error) {
	id := assign.GuideID
	assign.GuideID = ""
	err = m.config.Request("/v3/smartguide/guides/"+id+"/assign", assign, nil)
	return
}

// Query 服务人员查询
//
// https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/smartguide/chapter3_3.shtml
func (m SmartGuide) Query(query Query) (result QueryResult, err error) {
	params := utils.SortKey(utils.Struct2Map(query))
	err = m.config.Request("/v3/smartguide/guides?"+params, nil, &result)
	return
}
