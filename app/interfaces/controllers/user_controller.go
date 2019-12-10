package controllers

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/usecase"

	gin "github.com/gin-gonic/gin"
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
		c.JSON(403, NewError(err))
		return
	}
	c.JSON(200, gin.H{"token": token, "permission": permissionLevel})
}

func (controller *UserController) LogoutUser(c Context) {
	token := c.GetHeader("token")
	amountOfAffected, err := controller.Interactor.EndUserSession(token)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, amountOfAffected)
}

func (controller *UserController) Authenticate(c Context) {
	token := c.GetHeader("token")
	bool := controller.Interactor.CheckSessionValidity(token)
	if bool {
		c.JSON(200, "ACCESS GRANTED")
		return
	} else {
		// c.JSON(401, "ACCESS DENIED")
		c.AbortWithStatusJSON(401, gin.H{"Message": "Unauthorized"})
		return
	}
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

func (controller *UserController) StoreUsers(c Context) {
	users := domain.Users{}
	c.BindJSON(&users)
	amountOfStored, err := controller.Interactor.Store(users)
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
