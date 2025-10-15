package pkg

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

func AliPay(privateKey, appId, orderSn, totalPrice string) string {
	var client, err = alipay.New(appId, privateKey, false)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var p = alipay.TradeWapPay{}
	p.NotifyURL = "http://12b1981f.r17.cpolar.top/v1/alipay/notify"
	p.ReturnURL = "http://12b1981f.r17.cpolar.top/v1/orders/return"
	p.Subject = "闪送订单支付"
	p.OutTradeNo = orderSn
	p.TotalAmount = totalPrice
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	return url.String()
}
