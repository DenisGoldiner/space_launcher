package repo

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/DenisGoldiner/space_launcher/internal/entities"
)

// UserEntity represents the user DB entity.
type UserEntity struct {
	ID        string          `db:"id"`
	FirstName string          `db:"first_name"`
	LastName  string          `db:"last_name"`
	Gender    entities.Gender `db:"gender"`
	Birthday  time.Time       `db:"birthday"`
}

type UserRepo struct{}

func (ur UserRepo) GetAllUsers(ctx context.Context, dbExec sqlx.ExtContext) ([]entities.User, error) {
	getAllUsersQuery := `SELECT id, first_name, last_name, gender, birthday FROM "user"`

	rows, err := dbExec.QueryxContext(ctx, getAllUsersQuery)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var allUsers []entities.User

	for rows.Next() {
		var user UserEntity
		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}

		allUsers = append(allUsers, entities.User(user))
	}

	return allUsers, nil
}

func (ur UserRepo) SaveUser(ctx context.Context, dbExec sqlx.ExtContext, u entities.User) (entities.User, error) {
	saveUserQuery := `
		INSERT INTO "user" (first_name, last_name, gender, birthday)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT ON CONSTRAINT user_unique DO UPDATE SET first_name = EXCLUDED.first_name
		RETURNING id, first_name, last_name, gender, birthday`

	var savedUser UserEntity

	row := dbExec.QueryRowxContext(ctx, saveUserQuery, u.FirstName, u.LastName, u.Gender, u.Birthday)
	if err := row.StructScan(&savedUser); err != nil {
		return entities.User{}, err
	}

	return entities.User(savedUser), nil
}
