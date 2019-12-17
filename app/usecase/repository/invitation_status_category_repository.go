package repository

type InvitationStatusCategoryRepository interface {
	FindKeyIdByStatus(string) (int, error)
}
