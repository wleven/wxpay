// @Time : 2020/7/8 14:49
// @Author : 黑白配
// @File : index_test.go
// @PackageName:smartGuide
// @Description:

package smartGuide

import (
	"github.com/wleven/wxpay/global"
	"testing"
)

var smartGuide = New(&global.V3)

func TestSmartGuide_Query(t *testing.T) {
	if data, err := smartGuide.Query(Query{StoreID: 20774227}); err == nil {
		t.Log(data)
	} else {
		t.Error(err)
	}
}

func TestSmartGuide_Assign(t *testing.T) {
	if data, err := smartGuide.Assign(Assign{
		GuideID:    "",
		OutTradeNo: "",
	}); err == nil {
		t.Log(data)
	} else {
		t.Error(err)
	}
}

func TestSmartGuide_Register(t *testing.T) {
	if data, err := smartGuide.Register(Register{
		CorpID:  "",
		StoreID: 0,
		Name:    "",
		UserId:  "",
		Mobile:  "",
		QRCode:  "",
		Avatar:  "",
	}); err == nil {
		t.Log(data)
	} else {
		t.Error(err)
	}
}
