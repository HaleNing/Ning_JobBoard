package param

type Result struct {
	Code int64  `json:"code"`
	Data []byte `json:"data"`
	Msg  string `json:"msg"`
}
