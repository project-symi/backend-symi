package usecase

type UserAuthRepository interface {
	IssueToken(string, string) (string, error)
	RegisterToken(string, string) (int, error)
	GetPermissionName(string) (string, error)
	ValidateToken(string) (bool, error)
	RevokeToken(string) (int, error)
}
