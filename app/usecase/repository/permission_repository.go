package repository

import "project-symi-backend/app/domain"

type PermissionRepository interface {
	FindAll() (permissions domain.Permissions, err error)
}
