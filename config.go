package membership

import (
	"github.com/spurtcms/auth"
	role "github.com/spurtcms/team-roles"
	"gorm.io/gorm"
)

type Type string

const (
	Postgres Type = "postgres"
	Mysql    Type = "mysql"
)

type Config struct {
	AuthEnable       bool
	PermissionEnable bool
	DB               *gorm.DB
	Auth             *auth.Auth
	DataBaseType     Type
	Permissions      *role.PermissionConfig
}

type Membership struct {
	AuthEnable       bool
	PermissionEnable bool
	AuthFlg          bool
	PermissionFlg    bool
	DB               *gorm.DB
	Auth             *auth.Auth
	Permissions      *role.PermissionConfig
	UserId           int
	DataAccess       int
}
