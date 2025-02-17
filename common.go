package membership

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorAuth       = errors.New("auth enabled not initialised")
	ErrorPermission = errors.New("permissions enabled not initialised")
)

func AuthandPermission(membership *Membership) error {

	//check auth enable if enabled, use auth pkg otherwise it will return error
	if membership.AuthEnable && !membership.Auth.AuthFlg {

		return ErrorAuth
	}
	//check permission enable if enabled, use team-role pkg otherwise it will return error
	if membership.PermissionEnable && !membership.Auth.PermissionFlg {

		return ErrorPermission

	}

	return nil
}

// function to hash the secret passwords
func hashingPassword(pass string) string {

	passbyte, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err != nil {

		panic(err)

	}

	return string(passbyte)
}
