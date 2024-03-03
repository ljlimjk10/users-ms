package user_role_assoc_models

import (
	"github.com/go-pg/pg/v10"
	role_models "github.com/ljlimjk10/users-ms/models/role_models"
	user_models "github.com/ljlimjk10/users-ms/models/user_models"
)

type UserRoleAssociation struct {
	tableName  struct{}          `pg:"user_role_association"`
	UserRoleID int               `pg:"user_role_id, pk"`
	UserID     int               `pg:"user_id, notnull"`
	RoleID     int               `pg:"role_id, notnull"`
	User       *user_models.User `pg:"rel:has-one"`
	Role       *role_models.Role `pg:"rel:has-one"`
}

func CreateUserRoleAssoc(db *pg.DB, newUserRoleAssoc *UserRoleAssociation) error {
	_, err := db.Model(newUserRoleAssoc).Insert()
	return err
}
