package util

// 统一返回json结构
func Result(code int, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"data":    data,
		"message": "success",
	}
}
