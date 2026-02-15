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

type mockPropertyUsecase struct {
	usecase *mock.MockPropertyRepository
}

func TestPropertyUsecase_List(t *testing.T) {
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

	payload := entity.ListPayload{Search: "Property"}

	testCases := []struct {
		name               string
		payload            entity.ListPayload
		mockFn             func(mock *mockPropertyUsecase)
		expectedProperties []entity.Property
		expectedErr        error
	}{
		{
			name: "error: db connection error",
			mockFn: func(mocks *mockPropertyUsecase) {
				mocks.usecase.EXPECT().List(
					gomock.Any(), entity.ListPayload{},
				).Return([]entity.Property{}, errors.New("db connection error"))
			},
			expectedErr: errors.New("db connection error"),
		},
		{
			name:    "success: found with filter search",
			payload: payload,
			mockFn: func(mocks *mockPropertyUsecase) {
				mocks.usecase.EXPECT().List(
					gomock.Any(), payload,
				).Return(properties, nil)
			},
			expectedProperties: properties,
		},
		{
			name: "success: found without filter",
			mockFn: func(mocks *mockPropertyUsecase) {
				mocks.usecase.EXPECT().List(
					gomock.Any(), entity.ListPayload{},
				).Return(properties, nil)
			},
			expectedProperties: properties,
		},
	}

	ctrl := gomock.NewController(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockPropertyUsecase{
				usecase: mock.NewMockPropertyRepository(ctrl),
			}

			usecase := usecase.NewPropertyUsecase(mock.usecase)
			if tc.mockFn != nil {
				tc.mockFn(mock)
			}

			properties, err := usecase.List(context.Background(), tc.payload)
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedProperties, properties)
			}
		})
	}
}

func TestPropertyUsecase_GetByID(t *testing.T) {
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
		name             string
		id               string
		mockFn           func(mock *mockPropertyUsecase)
		expectedProperty entity.Property
		expectedErr      error
	}{
		{
			name: "error: db connection error",
			id:   "123",
			mockFn: func(mocks *mockPropertyUsecase) {
				mocks.usecase.EXPECT().GetByID(
					gomock.Any(), "123",
				).Return(entity.Property{}, errors.New("db connection error"))
			},
			expectedErr: errors.New("db connection error"),
		},
		{
			name: "success: found",
			id:   "123",
			mockFn: func(mocks *mockPropertyUsecase) {
				mocks.usecase.EXPECT().GetByID(
					gomock.Any(), "123",
				).Return(property, nil)
			},
			expectedProperty: property,
		},
	}

	ctrl := gomock.NewController(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockPropertyUsecase{
				usecase: mock.NewMockPropertyRepository(ctrl),
			}

			usecase := usecase.NewPropertyUsecase(mock.usecase)
			if tc.mockFn != nil {
				tc.mockFn(mock)
			}

			property, err := usecase.GetByID(context.Background(), tc.id)
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedProperty, property)
			}
		})
	}
}
