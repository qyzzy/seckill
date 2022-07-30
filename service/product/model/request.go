package model

type ProductRequest struct {
	Id int64 `json:"id"`
}

type RecommendRequest struct {
	Cursor int64 `json:"cursor"`
	Ps     int64 `json:"ps"`
}
