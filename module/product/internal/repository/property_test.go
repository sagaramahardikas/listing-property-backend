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

func TestPropertyRepository_List(t *testing.T) {
	testCases := []struct {
		name               string
		payload            entity.ListPayload
		expectedProperties []entity.Property
	}{
		{
			name:    "success: found with filter search",
			payload: entity.ListPayload{Search: "villa"},
			expectedProperties: []entity.Property{
				{
					ID:                 "123",
					Title:              "Test Villa",
					Price:              100000,
					Facilities:         []string{"Facility 1", "Facility 2"},
					Images:             []string{"image1.jpg", "image2.jpg"},
					Banner:             "banner1.jpg",
					Description:        "This is a test villa.",
					TermsAndConditions: "These are the terms and conditions.",
				},
			},
		},
		{
			name: "success: found without filter",
			expectedProperties: []entity.Property{
				{
					ID:                 "123",
					Title:              "Test Property",
					Price:              100000,
					Facilities:         []string{"Facility 1", "Facility 2"},
					Images:             []string{"image1.jpg", "image2.jpg"},
					Banner:             "banner1.jpg",
					Description:        "This is a test property.",
					TermsAndConditions: "These are the terms and conditions.",
				},
				{
					ID:                 "124",
					Title:              "Test Property 2",
					Price:              200000,
					Facilities:         []string{"Facility A", "Facility B"},
					Images:             []string{"imageA.jpg", "imageB.jpg"},
					Banner:             "banner2.jpg",
					Description:        "This is another test property.",
					TermsAndConditions: "These are the terms and conditions for property 2.",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.Nil(t, err)
			repo := repository.NewPropertyRepository(db)
			rows := mock.NewRows([]string{"id", "title", "description", "images", "facilities", "banner", "price", "terms_and_conditions"})
			for _, property := range tc.expectedProperties {
				rows.AddRow(
					property.ID,
					property.Title,
					property.Description,
					strings.Join(property.Images, ","),
					strings.Join(property.Facilities, ","),
					property.Banner,
					property.Price,
					property.TermsAndConditions,
				)
			}

			whereQuery := ""
			if tc.payload.Search != "" {
				whereQuery = " WHERE title LIKE ?"
			}

			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title, description, images, facilities, banner, price, terms_and_conditions FROM properties" + whereQuery)).
				WillReturnRows(rows)

			got, err := repo.List(context.Background(), tc.payload)
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedProperties, got)
		})
	}
}

func TestPropertyRepository_GetByID(t *testing.T) {
	testCases := []struct {
		name             string
		id               string
		dbGetError       error
		expectedProperty entity.Property
		expectedErr      error
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
			expectedProperty: entity.Property{
				ID:                 "123",
				Title:              "Test Property",
				Price:              100000,
				Facilities:         []string{"Facility 1", "Facility 2"},
				Images:             []string{"image1.jpg", "image2.jpg"},
				Banner:             "banner1.jpg",
				Description:        "This is a test property.",
				TermsAndConditions: "These are the terms and conditions.",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.Nil(t, err)
			repo := repository.NewPropertyRepository(db)
			rows := mock.NewRows([]string{"id", "title", "description", "images", "facilities", "banner", "price", "terms_and_conditions"})
			if tc.expectedErr == nil {
				rows.AddRow(
					tc.expectedProperty.ID,
					tc.expectedProperty.Title,
					tc.expectedProperty.Description,
					strings.Join(tc.expectedProperty.Images, ","),
					strings.Join(tc.expectedProperty.Facilities, ","),
					tc.expectedProperty.Banner,
					tc.expectedProperty.Price,
					tc.expectedProperty.TermsAndConditions,
				)
			}

			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title, description, images, facilities, banner, price, terms_and_conditions FROM properties WHERE id = ?")).
				WithArgs(tc.id).
				WillReturnRows(rows).
				WillReturnError(tc.dbGetError)

			got, err := repo.GetByID(context.Background(), tc.id)
			if err != nil {
				assert.Equal(t, tc.expectedErr, err)
				return
			}

			assert.Equal(t, tc.expectedProperty, got)
		})
	}
}
