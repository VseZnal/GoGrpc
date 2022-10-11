package models

type CreateReq struct {
	Slug     string `json:"slug"`
	ParentId string `json:"parent_Id"`
	UserId   string `json:"user_id"`
	PostId   string `json:"post_id"`
	Content  string `json:"content"`
	Status   string `json:"status"`
}

type CreateRes struct {
	Slug     string `json:"slug"`
	ParentId string `json:"parent_Id"`
	UserId   string `json:"user_id"`
	PostId   string `json:"post_id"`
	Content  string `json:"content"`
	Status   string `json:"status"`
}
