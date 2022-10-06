package models

type DeleteReq struct {
	Slug string `json:"slug"`
}

type DeleteRes struct {
	Status int32 `json:"status"`
}
