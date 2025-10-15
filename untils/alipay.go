package untils

import "github.com/Addiction-fei/public/pkg"

type Pay interface {
	AliPay(orderSn, totalPrice string) string
}
type AliPay struct {
}

func (a *AliPay) AliPay(privateKey, appId, orderSn, totalPrice string) string {
	pay := pkg.AliPay(privateKey, appId, orderSn, totalPrice)
	return pay
}
