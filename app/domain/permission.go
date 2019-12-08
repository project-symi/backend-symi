package domain

import "strings"

type Permission struct {
	Id   int
	Name string
}

type Permissions []Permission

func (permissions Permissions) PermissionToId(name string) (id int) {
	for _, permission := range permissions {
		if strings.ToLower(permission.Name) == strings.ToLower(name) {
			id = permission.Id
			return
		}
	}
	return
}
