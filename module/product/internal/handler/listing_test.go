package handler_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/dana/module/product/entity"
	"example.com/dana/module/product/internal/handler"
	"example.com/dana/module/product/internal/usecase/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type mockListingHandler struct {
	usecase *mock.MockListingUsecase
}

func TestListingHandler_List(t *testing.T) {
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
		name           string
		mockFn         func(mock *mockListingHandler)
		expectedResult string
	}{
		{
			name: "error: usecase error",
			mockFn: func(mocks *mockListingHandler) {
				mocks.usecase.EXPECT().List(
					gomock.Any(),
				).Return([]entity.Listing{}, errors.New("usecase error"))
			},
			expectedResult: "usecase error\n",
		},
		{
			name: "success: found",
			mockFn: func(mocks *mockListingHandler) {
				mocks.usecase.EXPECT().List(
					gomock.Any(),
				).Return(listings, nil)
			},
			expectedResult: "{\"listings\":[{\"id\":\"123\",\"banner\":\"\",\"title\":\"Test Listing\",\"description\":\"This is a test listing.\",\"images\":[\"image1.jpg\",\"image2.jpg\"],\"facilities\":[\"Facility 1\",\"Facility 2\"],\"price\":100000,\"terms_and_conditions\":\"These are the terms and conditions.\"},{\"id\":\"124\",\"banner\":\"\",\"title\":\"Test Listing 2\",\"description\":\"This is another test listing.\",\"images\":[\"imageA.jpg\",\"imageB.jpg\"],\"facilities\":[\"Facility A\",\"Facility B\"],\"price\":200000,\"terms_and_conditions\":\"These are the terms and conditions for listing 2.\"}]}\n",
		},
	}

	ctrl := gomock.NewController(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockListingHandler{
				usecase: mock.NewMockListingUsecase(ctrl),
			}

			r := httptest.NewRequest(http.MethodGet, "http://localhost/products/listings", nil)
			w := httptest.NewRecorder()
			handler := handler.NewListingHandler(mock.usecase)
			if tc.mockFn != nil {
				tc.mockFn(mock)
			}

			handler.ListListings()(w, r)
			body, _ := io.ReadAll(w.Body)
			responseText := string(body)
			assert.Equal(t, tc.expectedResult, responseText)
		})
	}
}

func TestListingHandler_GetListing(t *testing.T) {
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
		name           string
		id             string
		mockFn         func(mock *mockListingHandler)
		expectedResult string
	}{
		{
			name: "error: usecase error",
			id:   "123",
			mockFn: func(mocks *mockListingHandler) {
				mocks.usecase.EXPECT().GetByID(
					gomock.Any(), "123",
				).Return(entity.Listing{}, errors.New("usecase error"))
			},
			expectedResult: "usecase error\n",
		},
		{
			name: "success: found",
			id:   "123",
			mockFn: func(mocks *mockListingHandler) {
				mocks.usecase.EXPECT().GetByID(
					gomock.Any(), "123",
				).Return(listing, nil)
			},
			expectedResult: "{\"listing\":{\"id\":\"123\",\"banner\":\"\",\"title\":\"Test Listing\",\"description\":\"This is a test listing.\",\"images\":[\"image1.jpg\",\"image2.jpg\"],\"facilities\":[\"Facility 1\",\"Facility 2\"],\"price\":100000,\"terms_and_conditions\":\"These are the terms and conditions.\"}}\n",
		},
	}

	ctrl := gomock.NewController(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockListingHandler{
				usecase: mock.NewMockListingUsecase(ctrl),
			}

			r := httptest.NewRequest(http.MethodGet, "http://localhost/products/listings/123", nil)
			r.SetPathValue("id", tc.id)
			w := httptest.NewRecorder()
			handler := handler.NewListingHandler(mock.usecase)
			if tc.mockFn != nil {
				tc.mockFn(mock)
			}

			handler.GetListing()(w, r)
			body, _ := io.ReadAll(w.Body)
			responseText := string(body)
			assert.Equal(t, tc.expectedResult, responseText)
		})
	}
}
