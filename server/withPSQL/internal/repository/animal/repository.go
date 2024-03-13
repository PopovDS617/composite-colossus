package animal

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"withpsql/internal/client/db"
	"withpsql/internal/model"

	"withpsql/internal/repository"
)

const (
	tableName = "animal"

	idColumn        = "id"
	titleColumn     = "title"
	contentColumn   = "content"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.AnimalRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.Animal) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(titleColumn, contentColumn).
		// Values(info.Title, info.Content).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "animal_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Animal, error) {
	builder := sq.Select(idColumn, titleColumn, contentColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "animal_repository.Get",
		QueryRaw: query,
	}

	var animal model.Animal
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&animal.ID, &animal.Name, &animal.Type, &animal.Age, &animal.CreatedAt, &animal.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &animal, nil
}
