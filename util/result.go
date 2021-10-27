package util

const (
	SUCCESS = 0  // 成功
	ERROR   = -1 // 异常
)

// Result 统一返回json结构
func Result(code int, data interface{}, message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"data":    data,
		"message": message,
	}
}
