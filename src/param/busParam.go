package param

type BusCreateParam struct {
	BusName string `json:"busName"`
}

type BusUpdateParam struct {
	Id      int64  `json:"id"`
	BusName string `json:"busName"`
}
