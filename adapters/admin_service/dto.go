package admin_service

type AdminServiceRequest struct {
	Token string `json:"token"`
}

type AdminServiceResponse struct {
	Meta struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	IsValidToken bool               `json:"is_valid_token"`
	UserID       int64              `json:"user_id"`
	UserName     string             `json:"user_name"`
	Permissions  map[string][]int64 `json:"permissions"`
}
