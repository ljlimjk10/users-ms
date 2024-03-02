package user_models

import "time"

type User struct {
	tableName  struct{}  `pg:"users"`
	UserID     int       `pg:"user_id, pk"`
	Username   string    `pg:"username, notnull"`
	Email      string    `pg:"email, unique, notnull"`
	HashedPass string    `pg:"hashed_pass, notnull"`
	FirstName  string    `pg:"first_name"`
	LastName   string    `pg:"last_name"`
	CreatedAt  time.Time `pg:"created_at, notnull"`
	ModifiedAt time.Time `pg:"modified_at, notnull"`
}
