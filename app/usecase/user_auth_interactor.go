package usecase

type UserAuthInteractor struct {
	UserAuthRepository UserAuthRepository
}

func (interactor *UserAuthInteractor) CheckUserPass(employeeId string, employeePass string) (tokenId string, permissionLevel string, err error) {
	//GET THE PERMISSION LEVEL
	permissionLevel, err = interactor.UserAuthRepository.GetPermissionName(employeeId)
	if err != nil {
		return
	}

	//GENERATE TOCKEN ID IF EMPLOYEE INFO IS VALID
	tokenId, err = interactor.UserAuthRepository.IssueToken(employeeId, employeePass)
	if err != nil {
		return
	}

	//ADD THE GENERATED TOKEN ID TO THE USER TABLE
	amountOfAffected, err := interactor.UserAuthRepository.RegisterToken(employeeId, tokenId)
	if err != nil && amountOfAffected != 1 {
		return
	}

	return
}

func (interactor *UserAuthInteractor) CheckSessionValidity(tokenId string) (isValid bool, err error) {
	//CHECK IF RECEIVED ID IS VALID
	isValid, err = interactor.UserAuthRepository.ValidateToken(tokenId)
	return
}

func (interactor *UserAuthInteractor) EndUserSession(tokenId string) (amountOfDeleted int, err error) {
	amountOfDeleted, err = interactor.UserAuthRepository.RevokeToken(tokenId)
	return
}
