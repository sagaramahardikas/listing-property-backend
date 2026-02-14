package usecase

import (
	"context"

	"example.com/dana/module/product/entity"
	"example.com/dana/module/product/internal/repository"
)

type ListingUsecase interface {
	List(ctx context.Context) ([]entity.Listing, error)
	GetByID(ctx context.Context, id string) (entity.Listing, error)
}

type listingUsecase struct {
	listingRepo repository.ListingRepository
}

func (u *listingUsecase) List(ctx context.Context) ([]entity.Listing, error) {
	return u.listingRepo.List(ctx)
}

func (u *listingUsecase) GetByID(ctx context.Context, id string) (entity.Listing, error) {
	return u.listingRepo.GetByID(ctx, id)
}

func NewListingUsecase(listingRepo repository.ListingRepository) ListingUsecase {
	return &listingUsecase{
		listingRepo: listingRepo,
	}
}
