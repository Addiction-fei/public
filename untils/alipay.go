package untils

import "github.com/Addiction-fei/public/pkg"

type Pay interface {
	AliPay(orderSn, totalPrice string) string
}
type AliPay struct {
}

func (a *AliPay) AliPay(orderSn, totalPrice string) string {
	pay := pkg.AliPay(orderSn, totalPrice)
	return pay
}
