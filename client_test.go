package wxpay

import (
	"fmt"
	"testing"
)

func TestClient_UnifiedOrder(t *testing.T) {
	client := NewClient(NewAccount("xxxxx", "xxx", "xxxxx", false))
	params := make(Params)
	params.SetString("body", "test").
		SetString("out_trade_no", "58867657575757").
		SetInt64("total_fee", 1).
		SetString("spbill_create_ip", "127.0.0.1").
		SetString("notify_url", "http://notify.objcoding.com/notify").
		SetString("trade_type", "APP")
	t.Log(client.UnifiedOrder(params))
}

func TestClient_TransfersToUserDibByOpenid(t *testing.T) {
	// 设置支付账户
	payAccount := NewAccount("", "", "", false)
	// 设置支付证书
	payAccount.SetCertData("./xxxx.p12")
	// 加载支付终端
	client := NewClient(payAccount)
	// 加载请求参数
	params := make(Params)
	params.SetString("mch_appid", payAccount.appID) // 申请商户号的appid或商户号绑定的appid
	params.SetString("mchid", payAccount.mchID)     // 微信支付分配的商户号
	//params.SetString("device_info", "")               // 微信支付分配的终端设备号（非必填字段）
	params.SetString("nonce_str", NonceStr())     // 随机字符串，不长于32位
	params.SetString("partner_trade_no", "xxxxx") // 商户订单号 商户订单号，需保持唯一性 (只能是字母或者数字，不能包含有其它字符)
	params.SetString("openid", "xxxxxx")          // 商户appid下，某用户的openid
	params.SetString("check_name", "FORCE_CHECK") // 校验用户姓名选项(默认校验真实姓名) NO_CHECK：不校验真实姓名 FORCE_CHECK：强校验真实姓名
	params.SetString("re_user_name", "xxxxx")     // 收款用户姓名
	params.SetInt64("amount", 1)                  // 付款金额，单位为分
	params.SetString("desc", "测试付款")              // 付款备注(小于100字符)
	params.SetString("spbill_create_ip", "xxxxx") // Ip地址
	params.SetString("sign", client.Sign(params)) // 签名 - 签名算法:https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=4_3

	p, err := client.TransfersToUserDibByOpenid(params)
	fmt.Println(p, err)
}

func TestClient_TransfersGetTransferInfo(t *testing.T) {
	// 设置支付账户
	payAccount := NewAccount("", "", "", false)
	// 设置支付证书
	payAccount.SetCertData("./xxxx.p12")
	// 加载支付终端
	client := NewClient(payAccount)
	// 加载请求参数
	params := make(Params)
	params.SetString("appid", payAccount.appID)     // 申请商户号的appid或商户号绑定的appid
	params.SetString("mch_id", payAccount.mchID)    // 微信支付分配的商户号
	params.SetString("nonce_str", NonceStr())       // 随机字符串，不长于32位
	params.SetString("partner_trade_no", "xxxx001") // 商户订单号 商户订单号，需保持唯一性 (只能是字母或者数字，不能包含有其它字符)
	params.SetString("sign", client.Sign(params))   // 签名 - 签名算法:https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=4_3

	p, err := client.TransfersGetTransferInfo(params)
	fmt.Println(p, err)
}
