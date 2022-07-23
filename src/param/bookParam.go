package param

type BookParam struct {
	BookName string `json:"bookName"`
	Author   string `json:"author" validate:"required"`
}
