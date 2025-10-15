package untils

import "github.com/Addiction-fei/public/pkg"

type Pay interface {
	AliPay(orderSn, totalPrice string) string
}
type AliPay struct {
}

func (a *AliPay) AliPay(subject, returnURL, notifyURL, privateKey, appId, orderSn, totalPrice string) string {
	pay := pkg.AliPay(
		privateKey,
		appId,
		orderSn,
		totalPrice,
		subject,
		returnURL,
		notifyURL)
	return pay
}
