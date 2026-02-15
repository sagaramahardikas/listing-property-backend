package entity

type ListResponse struct {
	Listings []Listing `json:"listings"`
}

type GetResponse struct {
	Listing Listing `json:"listing"`
}

type Listing struct {
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
