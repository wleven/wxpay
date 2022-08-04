// @Time : 2020/7/9 11:14
// @Author : 黑白配
// @File : index
// @PackageName:public
// @Description:

package config

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/wleven/wxpay/utils"
)

type V3 struct {
	MchID     string `json:"mchid"`     // 商户ID
	ClientKey []byte `json:"clientKey"` // 证书私钥内容
	SerialNo  string `json:"serialNo"`  // 证书编号
}

func (m V3) Request(api string, body interface{}, out interface{}) error {

	var request *http.Request
	client := http.Client{}

	host := "https://api.mch.weixin.qq.com"
	if body != nil {
		b, _ := json.Marshal(body)
		request, _ = http.NewRequest("POST", host+api, bytes.NewReader(b))
	} else {
		request, _ = http.NewRequest("GET", host+api, nil)
	}
	sign, err := m.Sign(api, body)
	if err != nil {
		return err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", sign)
	if resp, err := client.Do(request); err == nil {
		defer resp.Body.Close()
		b, _ := ioutil.ReadAll(resp.Body)
		if len(b) > 0 && out != nil {
			return json.Unmarshal(b, out)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Sign 签名
// @param method string 请求类型
// @param url string 请求地址
// @param body interface 请求数据
func (m V3) Sign(url string, body interface{}) (string, error) {
	var data string = ""
	var method = "GET"
	t := time.Now().Unix()
	randomStr := utils.RandomStr()

	if body != nil {
		b, _ := json.Marshal(body)
		data = string(b)
		method = "POST"
	}

	str := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", method, url, t, randomStr, data)

	sign, err := m.rsaEncrypt([]byte(str), m.ClientKey)

	return fmt.Sprintf("WECHATPAY2-SHA256-RSA2048 mchid=\"%s\",nonce_str=\"%s\","+
		"signature=\"%s\",timestamp=\"%d\",serial_no=\"%s\"", m.MchID, randomStr, sign, t, m.SerialNo), err
}

// 私钥加密
func (m V3) rsaEncrypt(data, keyBytes []byte) (string, error) {
	//解密pem格式的私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return "", errors.New("public key error")
	}
	// 解析私钥
	pubInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PrivateKey)
	//加密
	msg := sha256.Sum256(data)
	sign, err := rsa.SignPKCS1v15(rand.Reader, pub, crypto.SHA256, msg[:])
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(sign), nil
}

type Result struct {
	Code    string `json:"code,omitempty"`    // 错误代码
	Message string `json:"message,omitempty"` // 错误信息
}
