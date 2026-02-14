package usecase_test

import (
	"context"
	"errors"
	"testing"

	"example.com/dana/module/product/entity"
	"example.com/dana/module/product/internal/repository/mock"
	"example.com/dana/module/product/internal/usecase"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type mockListingUsecase struct {
	usecase *mock.MockListingRepository
}

func TestListingUsecase_List(t *testing.T) {
	listings := []entity.Listing{
		{
			ID:                 "123",
			Title:              "Test Listing",
			Price:              100000,
			Facilities:         []string{"Facility 1", "Facility 2"},
			Images:             []string{"image1.jpg", "image2.jpg"},
			Description:        "This is a test listing.",
			TermsAndConditions: "These are the terms and conditions.",
		},
		{
			ID:                 "124",
			Title:              "Test Listing 2",
			Price:              200000,
			Facilities:         []string{"Facility A", "Facility B"},
			Images:             []string{"imageA.jpg", "imageB.jpg"},
			Description:        "This is another test listing.",
			TermsAndConditions: "These are the terms and conditions for listing 2.",
		},
	}

	testCases := []struct {
		name             string
		id               string
		mockFn           func(mock *mockListingUsecase)
		expectedListings []entity.Listing
		expectedErr      error
	}{
		{
			name: "error: db connection error",
			id:   "123",
			mockFn: func(mocks *mockListingUsecase) {
				mocks.usecase.EXPECT().List(
					gomock.Any(),
				).Return([]entity.Listing{}, errors.New("db connection error"))
			},
			expectedErr: errors.New("db connection error"),
		},
		{
			name: "success: found",
			id:   "123",
			mockFn: func(mocks *mockListingUsecase) {
				mocks.usecase.EXPECT().List(
					gomock.Any(),
				).Return(listings, nil)
			},
			expectedListings: listings,
		},
	}

	ctrl := gomock.NewController(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockListingUsecase{
				usecase: mock.NewMockListingRepository(ctrl),
			}

			usecase := usecase.NewListingUsecase(mock.usecase)
			if tc.mockFn != nil {
				tc.mockFn(mock)
			}

			listings, err := usecase.List(context.Background())
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedListings, listings)
			}
		})
	}
}

func TestListingUsecase_GetByID(t *testing.T) {
	listing := entity.Listing{
		ID:                 "123",
		Title:              "Test Listing",
		Price:              100000,
		Facilities:         []string{"Facility 1", "Facility 2"},
		Images:             []string{"image1.jpg", "image2.jpg"},
		Description:        "This is a test listing.",
		TermsAndConditions: "These are the terms and conditions.",
	}

	testCases := []struct {
		name            string
		id              string
		mockFn          func(mock *mockListingUsecase)
		expectedListing entity.Listing
		expectedErr     error
	}{
		{
			name: "error: db connection error",
			id:   "123",
			mockFn: func(mocks *mockListingUsecase) {
				mocks.usecase.EXPECT().GetByID(
					gomock.Any(), "123",
				).Return(entity.Listing{}, errors.New("db connection error"))
			},
			expectedErr: errors.New("db connection error"),
		},
		{
			name: "success: found",
			id:   "123",
			mockFn: func(mocks *mockListingUsecase) {
				mocks.usecase.EXPECT().GetByID(
					gomock.Any(), "123",
				).Return(listing, nil)
			},
			expectedListing: listing,
		},
	}

	ctrl := gomock.NewController(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockListingUsecase{
				usecase: mock.NewMockListingRepository(ctrl),
			}

			usecase := usecase.NewListingUsecase(mock.usecase)
			if tc.mockFn != nil {
				tc.mockFn(mock)
			}

			transaction, err := usecase.GetByID(context.Background(), tc.id)
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedListing, transaction)
			}
		})
	}
}
