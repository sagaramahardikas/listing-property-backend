package entity

type ListResponse struct {
	Properties []Property `json:"properties"`
}

type GetResponse struct {
	Property Property `json:"property"`
}

type Property struct {
	ID                 string   `json:"id"`
	Banner             string   `json:"banner"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Images             []string `json:"images"`
	Facilities         []string `json:"facilities"`
	Price              int      `json:"price"`
	TermsAndConditions string   `json:"terms_and_conditions"`
}

type ListPayload struct {
	Search string `json:"search"`
}
