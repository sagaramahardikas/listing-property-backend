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

type mockPropertyHandler struct {
	usecase *mock.MockPropertyUsecase
}

func TestPropertyHandler_List(t *testing.T) {
	properties := []entity.Property{
		{
			ID:                 "123",
			Title:              "Test Property",
			Price:              100000,
			Facilities:         []string{"Facility 1", "Facility 2"},
			Images:             []string{"image1.jpg", "image2.jpg"},
			Description:        "This is a test property.",
			TermsAndConditions: "These are the terms and conditions.",
		},
		{
			ID:                 "124",
			Title:              "Test Property 2",
			Price:              200000,
			Facilities:         []string{"Facility A", "Facility B"},
			Images:             []string{"imageA.jpg", "imageB.jpg"},
			Description:        "This is another test property.",
			TermsAndConditions: "These are the terms and conditions for property 2.",
		},
	}

	testCases := []struct {
		name           string
		payload        entity.ListPayload
		mockFn         func(mock *mockPropertyHandler)
		expectedResult string
	}{
		{
			name: "error: usecase error",
			mockFn: func(mocks *mockPropertyHandler) {
				mocks.usecase.EXPECT().List(
					gomock.Any(), entity.ListPayload{},
				).Return([]entity.Property{}, errors.New("usecase error"))
			},
			expectedResult: "usecase error\n",
		},
		{
			name:    "success: found",
			payload: entity.ListPayload{Search: "Property"},
			mockFn: func(mocks *mockPropertyHandler) {
				mocks.usecase.EXPECT().List(
					gomock.Any(), entity.ListPayload{Search: "Property"},
				).Return(properties, nil)
			},
			expectedResult: "{\"properties\":[{\"id\":\"123\",\"banner\":\"\",\"title\":\"Test Property\",\"description\":\"This is a test property.\",\"images\":[\"image1.jpg\",\"image2.jpg\"],\"facilities\":[\"Facility 1\",\"Facility 2\"],\"price\":100000,\"terms_and_conditions\":\"These are the terms and conditions.\"},{\"id\":\"124\",\"banner\":\"\",\"title\":\"Test Property 2\",\"description\":\"This is another test property.\",\"images\":[\"imageA.jpg\",\"imageB.jpg\"],\"facilities\":[\"Facility A\",\"Facility B\"],\"price\":200000,\"terms_and_conditions\":\"These are the terms and conditions for property 2.\"}]}\n",
		},
	}

	ctrl := gomock.NewController(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockPropertyHandler{
				usecase: mock.NewMockPropertyUsecase(ctrl),
			}

			queryParam := ""
			if tc.payload.Search != "" {
				queryParam = "?search=" + tc.payload.Search
			}

			r := httptest.NewRequest(http.MethodGet, "http://localhost/products/properties"+queryParam, nil)
			w := httptest.NewRecorder()
			handler := handler.NewPropertyHandler(mock.usecase)
			if tc.mockFn != nil {
				tc.mockFn(mock)
			}

			handler.ListProperties()(w, r)
			body, _ := io.ReadAll(w.Body)
			responseText := string(body)
			assert.Equal(t, tc.expectedResult, responseText)
		})
	}
}

func TestPropertyHandler_GetProperty(t *testing.T) {
	property := entity.Property{
		ID:                 "123",
		Title:              "Test Property",
		Price:              100000,
		Facilities:         []string{"Facility 1", "Facility 2"},
		Images:             []string{"image1.jpg", "image2.jpg"},
		Description:        "This is a test property.",
		TermsAndConditions: "These are the terms and conditions.",
	}

	testCases := []struct {
		name           string
		id             string
		mockFn         func(mock *mockPropertyHandler)
		expectedResult string
	}{
		{
			name: "error: usecase error",
			id:   "123",
			mockFn: func(mocks *mockPropertyHandler) {
				mocks.usecase.EXPECT().GetByID(
					gomock.Any(), "123",
				).Return(entity.Property{}, errors.New("usecase error"))
			},
			expectedResult: "usecase error\n",
		},
		{
			name: "success: found",
			id:   "123",
			mockFn: func(mocks *mockPropertyHandler) {
				mocks.usecase.EXPECT().GetByID(
					gomock.Any(), "123",
				).Return(property, nil)
			},
			expectedResult: "{\"property\":{\"id\":\"123\",\"banner\":\"\",\"title\":\"Test Property\",\"description\":\"This is a test property.\",\"images\":[\"image1.jpg\",\"image2.jpg\"],\"facilities\":[\"Facility 1\",\"Facility 2\"],\"price\":100000,\"terms_and_conditions\":\"These are the terms and conditions.\"}}\n",
		},
	}

	ctrl := gomock.NewController(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockPropertyHandler{
				usecase: mock.NewMockPropertyUsecase(ctrl),
			}

			r := httptest.NewRequest(http.MethodGet, "http://localhost/products/properties/123", nil)
			r.SetPathValue("id", tc.id)
			w := httptest.NewRecorder()
			handler := handler.NewPropertyHandler(mock.usecase)
			if tc.mockFn != nil {
				tc.mockFn(mock)
			}

			handler.GetProperty()(w, r)
			body, _ := io.ReadAll(w.Body)
			responseText := string(body)
			assert.Equal(t, tc.expectedResult, responseText)
		})
	}
}
