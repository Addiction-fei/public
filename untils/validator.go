package untils

import "regexp"

// 验证姓名格式
func IsValidRelaName(relaName string) bool {
	// 编译正则表达式
	reg := regexp.MustCompile(`^[a-zA-Z\p{Han}]{2,10}$`)
	// 使用正则表达式进行匹配
	return reg.MatchString(relaName)
}

// 验证身份证号格式
func IsValidIdCard(idCard string) bool {
	// 编译正则表达式
	reg := regexp.MustCompile(`^[1-9]\d{5}(18|19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[12]\d|3[01])\d{3}[\dXx]$`)
	// 使用正则表达式进行匹配
	return reg.MatchString(idCard)
}

// 验证手机号格式
func IsValidMobile(mobile string) bool {
	// 编译正则表达式
	reg := regexp.MustCompile(`^1[3-9]\d{9}$`)
	// 使用正则表达式进行匹配
	return reg.MatchString(mobile)
}

// 验证邮箱格式
func IsValidEmail(email string) bool {
	// 编译正则表达式
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	// 使用正则表达式进行匹配
	return re.MatchString(email)
}
