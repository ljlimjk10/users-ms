package user_role_assoc_models

import (
	role_models "github.com/ljlimjk10/users-ms/models/role_models"
	user_models "github.com/ljlimjk10/users-ms/models/user_models"
)

type UserRoleAssocation struct {
	tableName  struct{}          `pg:"user_role_association"`
	UserRoleID int               `pg:"user_role_id, pk"`
	UserID     int               `pg:"user_id, notnull"`
	RoleID     int               `pg:"role_id, notnull"`
	User       *user_models.User `pg:"rel:has-one"`
	Role       *role_models.Role `pg:"rel:has-one"`
}
