package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"strings"
	"testing"

	"example.com/dana/module/product/entity"
	"example.com/dana/module/product/internal/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestListingRepository_List(t *testing.T) {
	testCases := []struct {
		name             string
		expectedListings []entity.Listing
	}{
		{
			name: "success: found",
			expectedListings: []entity.Listing{
				{
					ID:                 "123",
					Title:              "Test Listing",
					Price:              100000,
					Facilities:         []string{"Facility 1", "Facility 2"},
					Images:             []string{"image1.jpg", "image2.jpg"},
					Banner:             "banner1.jpg",
					Description:        "This is a test listing.",
					TermsAndConditions: "These are the terms and conditions.",
				},
				{
					ID:                 "124",
					Title:              "Test Listing 2",
					Price:              200000,
					Facilities:         []string{"Facility A", "Facility B"},
					Images:             []string{"imageA.jpg", "imageB.jpg"},
					Banner:             "banner2.jpg",
					Description:        "This is another test listing.",
					TermsAndConditions: "These are the terms and conditions for listing 2.",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.Nil(t, err)
			repo := repository.NewListingRepository(db)
			rows := mock.NewRows([]string{"id", "title", "description", "images", "facilities", "banner", "price", "terms_and_conditions"})
			for _, listing := range tc.expectedListings {
				rows.AddRow(
					listing.ID,
					listing.Title,
					listing.Description,
					strings.Join(listing.Images, ","),
					strings.Join(listing.Facilities, ","),
					listing.Banner,
					listing.Price,
					listing.TermsAndConditions,
				)
			}

			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title, description, images, facilities, banner, price, terms_and_conditions FROM listings")).
				WillReturnRows(rows)

			got, err := repo.List(context.Background())
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedListings, got)
		})
	}
}

func TestListingRepository_GetByID(t *testing.T) {
	testCases := []struct {
		name            string
		id              string
		dbGetError      error
		expectedListing entity.Listing
		expectedErr     error
	}{
		{
			name:        "error: not found",
			id:          "123",
			dbGetError:  sql.ErrNoRows,
			expectedErr: sql.ErrNoRows,
		},
		{
			name:        "error: db connection error",
			id:          "123",
			dbGetError:  errors.New("db connection error"),
			expectedErr: errors.New("db connection error"),
		},
		{
			name: "success: found",
			id:   "123",
			expectedListing: entity.Listing{
				ID:                 "123",
				Title:              "Test Listing",
				Price:              100000,
				Facilities:         []string{"Facility 1", "Facility 2"},
				Images:             []string{"image1.jpg", "image2.jpg"},
				Banner:             "banner1.jpg",
				Description:        "This is a test listing.",
				TermsAndConditions: "These are the terms and conditions.",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.Nil(t, err)
			repo := repository.NewListingRepository(db)
			rows := mock.NewRows([]string{"id", "title", "description", "images", "facilities", "banner", "price", "terms_and_conditions"})
			if tc.expectedErr == nil {
				rows.AddRow(
					tc.expectedListing.ID,
					tc.expectedListing.Title,
					tc.expectedListing.Description,
					strings.Join(tc.expectedListing.Images, ","),
					strings.Join(tc.expectedListing.Facilities, ","),
					tc.expectedListing.Banner,
					tc.expectedListing.Price,
					tc.expectedListing.TermsAndConditions,
				)
			}

			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title, description, images, facilities, banner, price, terms_and_conditions FROM listings WHERE id = ?")).
				WithArgs(tc.id).
				WillReturnRows(rows).
				WillReturnError(tc.dbGetError)

			got, err := repo.GetByID(context.Background(), tc.id)
			if err != nil {
				assert.Equal(t, tc.expectedErr, err)
				return
			}

			assert.Equal(t, tc.expectedListing, got)
		})
	}
}
