package user_models

import (
	"time"

	"github.com/go-pg/pg/v10"
)

type User struct {
	tableName  struct{}  `pg:"users"`
	UserID     int       `pg:"user_id, pk"`
	Username   string    `pg:"username, notnull"`
	Email      string    `pg:"email, unique, notnull"`
	HashedPass string    `pg:"hashed_pass, notnull"`
	FirstName  string    `pg:"first_name"`
	LastName   string    `pg:"last_name"`
	CreatedAt  time.Time `pg:"created_at"`
	ModifiedAt time.Time `pg:"modified_at"`
}

func CreateUser(db *pg.DB, newUser *User) error {
	_, err := db.Model(newUser).Insert()
	return err
}
