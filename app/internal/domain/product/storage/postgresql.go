package storage

import (
	"production-snippets/internal/logging"

	"github.com/Masterminds/squirrel"
)

type ProductStorage struct {
	queryBuilder squirrel.StatementBuilderType
	client       PostgreSQLClient
	logger       *logging.Logger
}
