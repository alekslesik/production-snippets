package storage

import (
	"context"
	"production-snippets/internal/domain/product/model"
	"production-snippets/internal/logging"

	"github.com/Masterminds/squirrel"
)

type ProductStorage struct {
	queryBuilder squirrel.StatementBuilderType
	client       PostgreSQLClient
	logger       *logging.Logger
}

func NewProductStorage(client PostgreSQLClient, logger *logging.Logger) ProductStorage {
	squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	return ProductStorage{
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		client:       client,
		logger:       logger,
	}
}

const (
	scheme = "public"
	table = "product"
)

func (s *ProductStorage) All(ctx context.Context) ([]model.Prooduct, error) {
	// Select all colums
	query := s.queryBuilder.Select("id").
		Column("name").
		Column("descrition").
		Column("price").
		Column("currency_id").
		Column("image_id").
		Column("created_at").
		Column("updated_at").
		From(scheme + "." + table)

	sql, args, err := query.ToSql()
	if err != nil {
		s.logger.Err(err).Msg("Create query Error")
		return nil, err
	}

	rows, err := s.client.Query(ctx, sql, args...)
	if err != nil {
		s.logger.Err(err).Msg("Do query Error")
		return nil, err
	}

	defer rows.Close()

	list := make([]model.Prooduct, 0)

	for rows.Next() {
		p := model.Prooduct{}

		err := rows.Scan(
			&p.ID, 
			&p.Name, 
			&p.Descrition, 
			&p.Price, 
			&p.CurrencyId, 
			&p.ImageId, 
			&p.CreatedAt, 
			&p.UpdatedAt,
		)
		
		if err != nil {
			s.logger.Err(err).Msg("Scan Error")
			return nil, err
		}

		list = append(list, p)
	}

	return list, nil
}
