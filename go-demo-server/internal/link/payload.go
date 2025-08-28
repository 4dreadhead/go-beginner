package link

type CreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type UpdateRequest struct {
	Url  string `json:"url"   validate:"required,url"`
	Hash string `json:"hash"`
}

type LinksPageResponse struct {
	Links []Link `json:"links"`
	Count int64  `json:"count"`
}
