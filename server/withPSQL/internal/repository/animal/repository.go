package animal

import (
	"context"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"

	"withpsql/internal/client/db"
	"withpsql/internal/model"
	"withpsql/internal/repository/animal/converter"
	repoModel "withpsql/internal/repository/animal/model"

	"withpsql/internal/repository"
)

const (
	animalsTableName = "animals"
	regionsTableName = "regions"

	idColumn              = "id"
	nameColumn            = "name"
	typeColumn            = "type"
	genderColumn          = "gender"
	ageColumn             = "age"
	createdAtColumn       = "created_at"
	updatedAtColumn       = "updated_at"
	lastTimeSeenAtColumn  = "last_time_seen_at"
	animalsRegionIDColumn = "region_id"
	regionsRegionIDColumn = "region_id"
	seenByDeviceIDColumn  = "seen_by_device_id"
	regionNameColumn      = "region_name"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.AnimalRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, animal *model.Animal) (int64, error) {

	var regionID int

	getRegionIDBuilder := sq.Select(regionsRegionIDColumn).PlaceholderFormat(sq.Dollar).From(regionsTableName).Where(sq.Eq{regionNameColumn: animal.Region})

	query, args, err := getRegionIDBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "animal_repository.GetRegionID",
		QueryRaw: query,
	}

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&regionID)
	if err != nil {
		return 0, err
	}

	insertAnimalBuilder := sq.Insert(animalsTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, typeColumn, ageColumn, genderColumn, createdAtColumn, updatedAtColumn, animalsRegionIDColumn, lastTimeSeenAtColumn, seenByDeviceIDColumn).
		Values(animal.Name, animal.Type, animal.Age, animal.Gender, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339), regionID, time.Now().Format(time.RFC3339), 1).
		Suffix("RETURNING id")

	query, args, err = insertAnimalBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	q = db.Query{
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
	builder := sq.Select(idColumn, nameColumn, typeColumn, ageColumn, genderColumn, createdAtColumn, updatedAtColumn, fmt.Sprintf("%s.%s", animalsTableName, animalsRegionIDColumn), lastTimeSeenAtColumn, seenByDeviceIDColumn, regionNameColumn).
		PlaceholderFormat(sq.Dollar).
		From(animalsTableName).
		Where(sq.Eq{idColumn: id}).
		InnerJoin(fmt.Sprintf("%s ON %s.%s = %s.%s", regionsTableName, animalsTableName, animalsRegionIDColumn, regionsTableName, regionsRegionIDColumn)).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "animal_repository.Get",
		QueryRaw: query,
	}

	var animal repoModel.Animal
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&animal.ID,
		&animal.Name,
		&animal.Type,
		&animal.Age,
		&animal.Gender,
		&animal.CreatedAt,
		&animal.UpdatedAt,
		&animal.RegionID,
		&animal.LastTimeSeenAt,
		&animal.SeenByDevice,
		&animal.RegionName,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return converter.FromAnimalRepoToAnimalModel(&animal), nil
}

func (r *repo) GetAll(ctx context.Context) ([]*model.Animal, error) {
	builder := sq.Select(idColumn, nameColumn, typeColumn, ageColumn, genderColumn, createdAtColumn, updatedAtColumn, fmt.Sprintf("%s.%s", animalsTableName, animalsRegionIDColumn), lastTimeSeenAtColumn, seenByDeviceIDColumn, regionNameColumn).
		PlaceholderFormat(sq.Dollar).
		From(animalsTableName).
		InnerJoin(fmt.Sprintf("%s ON %s.%s = %s.%s", regionsTableName, animalsTableName, animalsRegionIDColumn, regionsTableName, regionsRegionIDColumn))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "animal_repository.GetAll",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var animals []*model.Animal
	for rows.Next() {
		var repoAnimal repoModel.Animal

		err := rows.Scan(&repoAnimal.ID,
			&repoAnimal.Name,
			&repoAnimal.Type,
			&repoAnimal.Age,
			&repoAnimal.Gender,
			&repoAnimal.CreatedAt,
			&repoAnimal.UpdatedAt,
			&repoAnimal.RegionID,
			&repoAnimal.LastTimeSeenAt,
			&repoAnimal.SeenByDevice,
			&repoAnimal.RegionName)

		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
		converted := converter.FromAnimalRepoToAnimalModel(&repoAnimal)

		animals = append(animals, converted)

	}

	return animals, nil
}
