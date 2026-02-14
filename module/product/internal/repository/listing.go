package repository

import (
	"context"
	"database/sql"
	"strings"

	"example.com/dana/module/product/entity"
	sq "github.com/Masterminds/squirrel"
)

type listingRepository struct {
	db *sql.DB
}

type ListingRepository interface {
	List(ctx context.Context) ([]entity.Listing, error)
	GetByID(ctx context.Context, id string) (entity.Listing, error)
}

func (r *listingRepository) List(ctx context.Context) ([]entity.Listing, error) {
	query := sq.Select(
		"id",
		"title",
		"description",
		"images",
		"facilities",
		"banner",
		"price",
		"terms_and_conditions",
	).From("listings")

	rows, err := query.RunWith(r.db).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	listings, err := scanListings(rows)
	if err != nil {
		return nil, err
	}

	return listings, nil
}

func (r *listingRepository) GetByID(ctx context.Context, id string) (entity.Listing, error) {
	query := sq.Select(
		"id",
		"title",
		"description",
		"images",
		"facilities",
		"banner",
		"price",
		"terms_and_conditions",
	).From("listings").Where(sq.Eq{"id": id})

	row := query.RunWith(r.db).QueryRowContext(ctx)
	listing, err := scanListing(row)
	if err != nil {
		return entity.Listing{}, err
	}

	return listing, nil
}

func NewListingRepository(db *sql.DB) ListingRepository {
	return &listingRepository{
		db: db,
	}
}

func parseStringList(value string) []string {
	if strings.TrimSpace(value) == "" {
		return []string{}
	}

	return strings.Split(value, ",")
}

func scanListings(rows *sql.Rows) ([]entity.Listing, error) {
	var listings []entity.Listing

	for rows.Next() {
		var listing entity.Listing
		var images sql.NullString
		var facilities sql.NullString

		if err := rows.Scan(
			&listing.ID,
			&listing.Title,
			&listing.Description,
			&images,
			&facilities,
			&listing.Banner,
			&listing.Price,
			&listing.TermsAndConditions,
		); err != nil {
			return nil, err
		}

		listing.Images = parseStringList(images.String)
		listing.Facilities = parseStringList(facilities.String)

		listings = append(listings, listing)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listings, nil
}

func scanListing(row sq.RowScanner) (entity.Listing, error) {
	var listing entity.Listing
	var images sql.NullString
	var facilities sql.NullString

	if err := row.Scan(
		&listing.ID,
		&listing.Title,
		&listing.Description,
		&images,
		&facilities,
		&listing.Banner,
		&listing.Price,
		&listing.TermsAndConditions,
	); err != nil {
		if err == sql.ErrNoRows {
			return entity.Listing{}, nil
		}

		return entity.Listing{}, err
	}

	listing.Images = parseStringList(images.String)
	listing.Facilities = parseStringList(facilities.String)

	return listing, nil
}
