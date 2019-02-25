package e

var MsgFlag = map[int]string{
	SUCCESS : "ok",
	ERROR : "error",
	INVALID_ERROR : "未定义的错误",
	INVALID_PARAMS : "请求参数错误",
	ERROR_TOKEN_FAIL : "错误的token",
	ERROR_TOKEN_TIME_OUT : "token过期",

	ERROR_AUTH_TOKEN : "Token生成失败",
	ERROR_AUTH : "错误的用户信息",

	TAGS_IS_EXIST : "这个标签已经存在",
	TAGS_INSERT_FAIL : "标签添加失败",

	ERROR_UPLOAD_CHECK_IMAGE_FORMAT : "校验图片错误，图片格式或大小有问题",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL : "检查图片失败",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL : "保存图片失败",
}

func GetMsg(code int) string{
	if msg,ok := MsgFlag[code]; ok{
		return msg
	}
	return MsgFlag[INVALID_ERROR]
}
