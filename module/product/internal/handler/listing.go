package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/dana/module/product/entity"
)

type ListingHandler struct {
	//usecase usecase.CommerceTransaction
	logger *log.Logger
}

func (h *ListingHandler) ListListings() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := entity.ListListingResponse{
			Data: entity.DataListListing{
				Listings: []entity.Listing{
					{
						ID:          "1",
						Banner:      "https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg",
						Title:       "My Villa Sample",
						Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
						Images: []string{
							"https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg",
							"https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg",
							"https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg",
							"https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg",
						},
						Facilities:         []string{"Kitchen", "Free Park", "Bar"},
						Price:              750000,
						TermsAndConditions: "At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias excepturi sint occaecati cupiditate non provident, similique sunt in culpa qui officia deserunt mollitia animi, id est laborum et dolorum fuga. Et harum quidem rerum facilis est et expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio cumque nihil impedit quo minus id quod maxime placeat facere possimus, omnis voluptas assumenda est, omnis dolor repellendus. Temporibus autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet ut et voluptates repudiandae sint et molestiae non recusandae. Itaque earum rerum hic tenetur a sapiente delectus, ut aut reiciendis voluptatibus maiores alias consequatur aut perferendis doloribus asperiores repellat.",
					},
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *ListingHandler) GetListing() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := entity.GetListingResponse{
			Data: entity.DataGetListing{
				Listing: entity.Listing{
					ID:          "1",
					Banner:      "https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg",
					Title:       "My Villa Sample",
					Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
					Images: []string{
						"https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg",
						"https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg",
						"https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg",
						"https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg",
					},
					Facilities:         []string{"Kitchen", "Free Park", "Bar"},
					Price:              750000,
					TermsAndConditions: "At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias excepturi sint occaecati cupiditate non provident, similique sunt in culpa qui officia deserunt mollitia animi, id est laborum et dolorum fuga. Et harum quidem rerum facilis est et expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio cumque nihil impedit quo minus id quod maxime placeat facere possimus, omnis voluptas assumenda est, omnis dolor repellendus. Temporibus autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet ut et voluptates repudiandae sint et molestiae non recusandae. Itaque earum rerum hic tenetur a sapiente delectus, ut aut reiciendis voluptatibus maiores alias consequatur aut perferendis doloribus asperiores repellat.",
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func NewListingHandler() *ListingHandler {
	return &ListingHandler{}
}
