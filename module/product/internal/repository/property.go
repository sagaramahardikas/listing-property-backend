package repository

import (
	"context"
	"database/sql"
	"strings"

	"example.com/dana/module/product/entity"
	sq "github.com/Masterminds/squirrel"
)

type propertyRepository struct {
	db *sql.DB
}

type PropertyRepository interface {
	List(ctx context.Context, payload entity.ListPayload) ([]entity.Property, error)
	GetByID(ctx context.Context, id string) (entity.Property, error)
}

func (r *propertyRepository) List(ctx context.Context, payload entity.ListPayload) ([]entity.Property, error) {
	query := sq.Select(
		"id",
		"title",
		"description",
		"images",
		"facilities",
		"banner",
		"price",
		"terms_and_conditions",
	).From("properties")

	if payload.Search != "" {
		query = query.Where(sq.Like{"title": "%" + payload.Search + "%"})
	}

	rows, err := query.RunWith(r.db).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	properties, err := scanProperties(rows)
	if err != nil {
		return nil, err
	}

	return properties, nil
}

func (r *propertyRepository) GetByID(ctx context.Context, id string) (entity.Property, error) {
	query := sq.Select(
		"id",
		"title",
		"description",
		"images",
		"facilities",
		"banner",
		"price",
		"terms_and_conditions",
	).From("properties").Where(sq.Eq{"id": id})

	row := query.RunWith(r.db).QueryRowContext(ctx)
	property, err := scanProperty(row)
	if err != nil {
		return entity.Property{}, err
	}

	return property, nil
}

func NewPropertyRepository(db *sql.DB) PropertyRepository {
	return &propertyRepository{
		db: db,
	}
}

func parseStringList(value string) []string {
	if strings.TrimSpace(value) == "" {
		return []string{}
	}

	return strings.Split(value, ",")
}

func scanProperties(rows *sql.Rows) ([]entity.Property, error) {
	properties := make([]entity.Property, 0)

	for rows.Next() {
		var property entity.Property
		var images sql.NullString
		var facilities sql.NullString

		if err := rows.Scan(
			&property.ID,
			&property.Title,
			&property.Description,
			&images,
			&facilities,
			&property.Banner,
			&property.Price,
			&property.TermsAndConditions,
		); err != nil {
			return nil, err
		}

		property.Images = parseStringList(images.String)
		property.Facilities = parseStringList(facilities.String)

		properties = append(properties, property)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return properties, nil
}

func scanProperty(row sq.RowScanner) (entity.Property, error) {
	var property entity.Property
	var images sql.NullString
	var facilities sql.NullString

	if err := row.Scan(
		&property.ID,
		&property.Title,
		&property.Description,
		&images,
		&facilities,
		&property.Banner,
		&property.Price,
		&property.TermsAndConditions,
	); err != nil {
		if err == sql.ErrNoRows {
			return entity.Property{}, nil
		}

		return entity.Property{}, err
	}

	property.Images = parseStringList(images.String)
	property.Facilities = parseStringList(facilities.String)

	return property, nil
}
