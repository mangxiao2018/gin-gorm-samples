package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_EXIST_USER:                "已存在该用户名称",
	ERROR_EXIST_USER_FAIL:           "获取已存在用户失败",
	ERROR_NOT_EXIST_USER:            "该用户不存在",
	ERROR_GET_USERS_FAIL:            "获取所有用户失败",
	ERROR_COUNT_USER_FAIL:           "统计用户失败",
	ERROR_ADD_USER_FAIL:             "新增用户失败",
	ERROR_EDIT_USER_FAIL:            "修改用户失败",
	ERROR_DELETE_USER_FAIL:          "删除用户失败",
	ERROR_EXPORT_USER_FAIL:          "导出用户失败",
	ERROR_IMPORT_USER_FAIL:          "导入用户失败",
	ERROR_NOT_EXIST_ORDER:           "该订单不存在",
	ERROR_ADD_ORDER_FAIL:            "新增订单失败",
	ERROR_DELETE_ORDER_FAIL:         "删除订单失败",
	ERROR_CHECK_EXIST_ORDER_FAIL:    "检查订单是否存在失败",
	ERROR_EDIT_ORDER_FAIL:           "修改订单失败",
	ERROR_COUNT_ORDER_FAIL:          "统计订单失败",
	ERROR_GET_ORDERS_FAIL:           "获取多个订单失败",
	ERROR_GET_ORDER_FAIL:            "获取单个订单失败",
	ERROR_GEN_ORDER_POSTER_FAIL:     "生成订单失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
