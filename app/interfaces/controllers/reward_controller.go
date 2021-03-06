package controllers

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/usecase/interactor"
)

type RewardController struct {
	Interactor interactor.RewardInteractor
}

func NewRewardController(sqlHandler database.SqlHandler) *RewardController {
	return &RewardController{
		Interactor: interactor.RewardInteractor{
			RewardRepository: &database.RewardRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *RewardController) AllRewards(c Context) {
	rewards, err := controller.Interactor.FindAllRewards()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, rewards)
}

func (controller *RewardController) PatchRewardById(c Context) {
	reward := domain.Reward{}
	if err := c.ShouldBindJSON(&reward); err != nil {
		c.JSON(400, ValidationError("PatchRewardById method's json parameter is invalid ", err))
		return
	}
	rewards, err := controller.Interactor.PatchById(reward)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	if rewards == false {
		c.Status(400)
		return
	}
	c.Status(200)
}
