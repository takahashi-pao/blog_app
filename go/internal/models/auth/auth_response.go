package auth_response_model

// 認証系レスポンスモデル
type Auth_Response_Props struct {
	Message  string `json:"message"`
	Error    string `json:"error"`
	UserId   string `json:"userId"`
	IsSignIn bool   `json:"isSignIn"`
}
