package entity

type ListListingResponse struct {
	Data DataListListing `json:"data"`
}

type GetListingResponse struct {
	Data DataGetListing `json:"data"`
}

type DataListListing struct {
	Listings []Listing `json:"listings"`
}

type DataGetListing struct {
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
