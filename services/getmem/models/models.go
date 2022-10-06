package models

type GetReq struct {
	Slug string `json:"slug"`
}

type GetRes struct {
	Slug     string `json:"slug"`
	ParentId string `json:"parent_Id"`
	UserId   string `json:"user_id"`
	PostId   string `json:"post_id"`
	Content  string `json:"content"`
	Status   string `json:"status"`
}
