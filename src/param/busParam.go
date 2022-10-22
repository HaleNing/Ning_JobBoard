package param

type JobCreateParam struct {
	CompanyName string `json:"company_name" `
	JobExp      int8   `json:"job_exp"`
	JobDesc     string `json:"job_desc"`
	JobName     string `json:"job_name"`
	IsRemote    bool   `json:"is_remote"`
	JobArea     string `json:"job_area"`
}

type BusUpdateParam struct {
	Id      int64  `json:"id"`
	BusName string `json:"busName"`
}
