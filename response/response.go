package response

import "github.com/gin-gonic/gin"

func SuccessResponse(data ...interface{}) gin.H {
	code := 0
	message := "成功"
	params := gin.H{}

	for _, value := range data {
		switch value.(type) {
		case int:
			code = value.(int)
		case string:
			message = value.(string)
		case gin.H:
			params = value.(gin.H)
		}
	}

	return Response(code, message, params)
}

func ErrorResponse(data ...interface{}) gin.H {
	code := -1
	message := "错误"
	params := gin.H{}

	for _, value := range data {
		switch value.(type) {
		case int:
			code = value.(int)
		case string:
			message = value.(string)
		case gin.H:
			params = value.(gin.H)
		}
	}

	return Response(code, message, params)
}

func Response(code int, message string, params gin.H) gin.H {
	response := gin.H{
		"code":    code,
		"message": message,
	}

	for index, value := range params {
		response[index] = value
	}

	return response
}
