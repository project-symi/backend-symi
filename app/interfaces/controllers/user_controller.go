package controllers

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/usecase/interactor"
)

type UserController struct {
	Interactor interactor.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: interactor.UserInteractor{
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

func (controller *UserController) AllUsers(c Context) {
	userRoot, err := controller.Interactor.Users()
	users := userRoot.Users
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	if err = c.ShouldBind(&userRoot); err != nil {
		info := "AllUsers method: "
		c.JSON(400, ValidationError(info, err))
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) TopPointsUsers(c Context) {
	const NumOfRank = 7
	users, err := controller.Interactor.FindTopPointsUsers(NumOfRank)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) UserByEmployeeId(c Context) {
	employeeId := domain.EmployeeIdParam{}
	if err := c.ShouldBindUri(&employeeId); err != nil {
		c.JSON(400, ValidationError("UsersByEmployeeId method's parameter is invalid ", err))
		return
	}
	user, err := controller.Interactor.User(employeeId.EmployeeId)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
}

func (controller *UserController) UsersByEmployeeName(c Context) {
	name := domain.NameQuery{}
	if err := c.ShouldBind(&name); err != nil {
		c.JSON(400, ValidationError("UsersByEmployeeName method's query string is invalid ", err))
		return
	}
	users, err := controller.Interactor.UsersByName(name.Name)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) DeleteByEmployeeId(c Context) {
	employeeId := domain.EmployeeIdParam{}
	if err := c.ShouldBindUri(&employeeId); err != nil {
		c.JSON(400, ValidationError("DeleteByEmployeeId method's parameter is invalid ", err))
		return
	}
	amountOfDeleted, err := controller.Interactor.Delete(employeeId.EmployeeId)
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
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, ValidationError("StoreUser method", err))
		return
	}
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
	root := domain.UsersRoot{}
	if err := c.ShouldBindJSON(&root); err != nil {
		info := "StoreUsers: "
		c.JSON(400, ValidationError(info, err))
		return
	}
	amountOfStored, err := controller.Interactor.StoreUsers(root)
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
