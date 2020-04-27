package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Struct2Map 结构体转map
func Struct2Map(params interface{}) (m map[string]interface{}) {
	m = make(map[string]interface{})

	if b, err := json.Marshal(params); err == nil {
		_ = json.Unmarshal(b, &m)
	}
	return
}

// SortKey map排序
func SortKey(m map[string]interface{}) string {
	arr := make([]string, 0)
	str := ""
	for k := range m {
		arr = append(arr, k)
	}
	sort.Strings(arr)
	for _, v := range arr {
		switch m[v].(type) {
		case string:
			str = str + v + "=" + m[v].(string) + "&"
		case int:
			num := strconv.Itoa(m[v].(int))
			str = str + v + "=" + num + "&"
		}
	}
	return str
}

// MAP2XML map转xml
func MAP2XML(m map[string]interface{}) string {
	str := ""
	for k, v := range m {
		switch v.(type) {
		case string:
			str = str + fmt.Sprintf("<%s><![CDATA[%s]]></%s>", k, v, k)
		case int:
			str = str + fmt.Sprintf("<%s><![CDATA[%d]]></%s>", k, v, k)
		case interface{}:
			b, _ := json.Marshal(v)
			str = str + fmt.Sprintf("<%s><![CDATA[%s]]></%s>", k, string(b), k)
		}
	}
	return "<xml>" + str + "</xml>"
}

// MAPMerge map合并
func MAPMerge(args ...map[string]interface{}) (m map[string]interface{}) {
	m = make(map[string]interface{})
	for _, item := range args {
		for k, v := range item {
			m[k] = v
		}
	}
	return m
}

// XML2MAP xml转map
func XML2MAP(b []byte) (m map[string]string) {

	decoder := xml.NewDecoder(bytes.NewReader(b))
	m = make(map[string]string)
	tag := ""
	for {
		token, err := decoder.Token()

		if err != nil {
			break
		}
		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local != "xml" {
				tag = t.Name.Local
			} else {
				tag = ""
			}
			break
		case xml.EndElement:
			break
		case xml.CharData:
			data := strings.TrimSpace(string(t))
			if len(data) != 0 {
				m[tag] = data
			}
			break
		}
	}
	return
}

// RandomStr 随机字符串
func RandomStr() string {
	return strings.ToUpper(MD5(strconv.FormatInt(time.Now().UnixNano(), 19)))
}

// MD5 md5加密
func MD5(str string) string {
	hash := md5.Sum([]byte(str))
	md5str := fmt.Sprintf("%x", hash)
	return strings.ToUpper(md5str)
}

// Sign 生成签名 HMAC-SHA256加密
func Sign(m map[string]interface{}, key string) (sign string) {
	sign = HmacSha256(SortKey(m)+"key="+key, key)
	return
}

// HmacSha256 HMAC-SHA256加密
func HmacSha256(str string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(str))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
