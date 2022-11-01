package param

type JobSelectParam struct {
	Name     string `query:"name"`
	Tech     string `query:"tech"`
	IsRemote int8   `query:"remote"`
	Area     string `query:"area"`
}
