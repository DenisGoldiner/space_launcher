package repo

import (
	"context"
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/jmoiron/sqlx"
	"time"
)

type User struct {
	ID        string          `db:"id"`
	FirstName string          `db:"first_name"`
	LastName  string          `db:"last_name"`
	Gender    entities.Gender `db:"gender"`
	Birthday  time.Time       `db:"birthday"`
}

type UserRepo struct{}

func (ur UserRepo) SaveUser(ctx context.Context, dbExec sqlx.ExtContext, u entities.User) (entities.User, error) {
	saveUserQuery := `
		INSERT INTO "user" (first_name, last_name, gender, birthday)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT ON CONSTRAINT user_unique DO UPDATE SET first_name = EXCLUDED.first_name
		RETURNING (id, first_name, last_name, gender, birthday)`

	var savedUser User

	row := dbExec.QueryRowxContext(ctx, saveUserQuery, u.FirstName, u.LastName, u.Gender, u.Birthday)
	if err := row.StructScan(&savedUser); err != nil {
		return entities.User{}, err
	}

	return entities.User(savedUser), nil
}
