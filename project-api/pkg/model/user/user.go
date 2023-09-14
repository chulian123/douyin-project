package user

type RegisterReq struct {
	// 密码，最长32个字符
	Password string `json:"password"`
	// 注册用户名，最长32个字符
	Username string `json:"username"`
}

type RegisterRsp struct {
	// 状态码，0-成功，其他值-失败
	StatusCode int64 `json:"status_code"`
	// 返回状态描述
	StatusMsg string `json:"status_msg"`
	// 用户鉴权token
	Token string `json:"token"`
	// 用户id
	UserID int64 `json:"user_id"`
}

type LoginReq struct {
	// 登录密码
	Password string `json:"password"`
	// 登录用户名
	Username string `json:"username"`
}

type LoginRsp struct {
	// 状态码，0-成功，其他值-失败
	StatusCode int64 `json:"status_code"`
	// 返回状态描述
	StatusMsg *string `json:"status_msg"`
	// 用户id
	UserID *int64 `json:"user_id"`
	// 用户鉴权token
	Token *string `json:"token"`
}
