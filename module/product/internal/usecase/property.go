package usecase

import (
	"context"

	"example.com/dana/module/product/entity"
	"example.com/dana/module/product/internal/repository"
)

type PropertyUsecase interface {
	List(ctx context.Context, payload entity.ListPayload) ([]entity.Property, error)
	GetByID(ctx context.Context, id string) (entity.Property, error)
}

type propertyUsecase struct {
	propertyRepo repository.PropertyRepository
}

func (u *propertyUsecase) List(ctx context.Context, payload entity.ListPayload) ([]entity.Property, error) {
	return u.propertyRepo.List(ctx, payload)
}

func (u *propertyUsecase) GetByID(ctx context.Context, id string) (entity.Property, error) {
	return u.propertyRepo.GetByID(ctx, id)
}

func NewPropertyUsecase(propertyRepo repository.PropertyRepository) PropertyUsecase {
	return &propertyUsecase{
		propertyRepo: propertyRepo,
	}
}
