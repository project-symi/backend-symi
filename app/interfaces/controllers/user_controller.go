package controllers

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
			GenderRepository: &database.GenderRepository{
				SqlHandler: sqlHandler,
			},
			DepartmentRepository: &database.DepartmentRepository{
				SqlHandler: sqlHandler,
			},
			PermissionRepository: &database.PermissionRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) LoginUser(c Context) {
	type UserCredentials struct {
		EmployeeId string `json:"userId"`
		Pass       string `json:"password"`
	}
	var user UserCredentials
	if err := c.BindJSON(&user); err != nil {
		panic(err)
	}

	token, permissionLevel, err := controller.Interactor.CheckUserPass(user.EmployeeId, user.Pass)
	if err != nil {
		c.JSON(401, NewError(err))
		return
	}

	type TokenResponse struct {
		Token           string `json:"token"`
		PermissionLevel string `json:"permission"`
	}

	response := TokenResponse{Token: token, PermissionLevel: permissionLevel}

	c.JSON(200, response)
}

func (controller *UserController) LogoutUser(c Context) {
	token := c.GetHeader("token")
	amountOfDeleted, err := controller.Interactor.EndUserSession(token)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	if amountOfDeleted == 0 {
		c.JSON(400, NewError(err)) //TODO: create another error
		return
	}
	c.JSON(200, amountOfDeleted)
}

func (controller *UserController) Authenticate(c Context) {
	token := c.GetHeader("token")
	tokenValid, err := controller.Interactor.CheckSessionValidity(token)
	if !tokenValid || err != nil {
		c.AbortWithStatus(401)
		return
	}
	return
}

func (controller *UserController) AllUsers(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) UserByEmployeeId(c Context) {
	user, err := controller.Interactor.User(c.Param("employeeId"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
}

func (controller *UserController) UsersByEmployeeName(c Context) {
	users, err := controller.Interactor.UsersByName(c.Query("name"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) DeleteByEmployeeId(c Context) {
	amountOfDeleted, err := controller.Interactor.Delete(c.Param("employeeId"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	if amountOfDeleted == 0 {
		c.JSON(400, NewError(err)) //TODO: create another error
		return
	}
	c.Status(204)
}

func (controller *UserController) StoreUser(c Context) {
	user := domain.User{}
	c.BindJSON(&user)
	success, err := controller.Interactor.StoreUser(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	if success == false {
		c.Status(400)
		return
	}
	c.Status(201)
}

func (controller *UserController) StoreUsers(c Context) {
	users := domain.Users{}
	c.BindJSON(&users)
	amountOfStored, err := controller.Interactor.StoreUsers(users)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	if amountOfStored == 0 {
		c.Status(200)
		return
	}
	c.Status(201)
}
