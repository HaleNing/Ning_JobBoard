package param

type UserLoginParam struct {
	UserName string `json:"user_name" validate:"required"`
	UserPass string `json:"user_pass" validate:"required"`
}
