package response

type (
	BaseResponse struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Result  interface{} `json:"result"`
	}

	JwtResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	User struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}

	Error struct {
		Error  string `json:"error"`
		Status int    `json:"status"`
	}
)
