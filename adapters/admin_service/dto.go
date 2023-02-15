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
	Active   bool     `json:"active"`
	Email    string   `json:"email"`
	UserID   string   `json:"user_id"`
	Tenant   string   `json:"tenant"`
	UserName string   `json:"username"`
	Scope    []string `json:"scope"`
}
