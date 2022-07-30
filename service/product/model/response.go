package model

type ProductResponse struct {
	Product Product `json:"product"`
}

type RecommendResponse struct {
	Products []*Product `json:"products"`
	IsEnd    bool       `json:"is_end"`
}
