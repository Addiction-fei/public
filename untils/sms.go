package untils

import "github.com/Addiction-fei/public/pkg"

type Sms interface {
	SendSms(phone string, code string) error
}
type SendSms struct {
}

func (s *SendSms) SendSms(signName, templateCode, phone, code string) error {
	sms, err := pkg.SendSms(signName, templateCode, phone, code)
	if err != nil {

		return err
	}
	if *sms.Body.Code != "OK" {
		return err
	}
	return nil
}
